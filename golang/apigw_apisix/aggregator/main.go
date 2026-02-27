package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type ClientInfo struct {
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
	Tier     string `json:"tier"`
}

type PriceResponse struct {
	ProductID  string  `json:"product_id"`
	ClientID   string  `json:"client_id"`
	ClientName string  `json:"client_name"`
	Tier       string  `json:"tier"`
	Price      float64 `json:"price"`
	Currency   string  `json:"currency"`
	Source     string  `json:"source"`
	PromoName  string  `json:"promo_name,omitempty"`
}

type serviceResponse struct {
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	PromoName string  `json:"promo_name,omitempty"`
}

var httpClient = &http.Client{Timeout: 5 * time.Second}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	if productID == "" {
		http.Error(w, `{"error":"product_id is required"}`, http.StatusBadRequest)
		return
	}

	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		http.Error(w, `{"error":"client_id is required"}`, http.StatusBadRequest)
		return
	}

	clientServiceURL := os.Getenv("CLIENT_SERVICE_URL")
	clientInfo, err := fetchClient(fmt.Sprintf("%s/client?client_id=%s", clientServiceURL, clientID))
	if err != nil {
		http.Error(w, `{"error":"client not found"}`, http.StatusNotFound)
		return
	}

	promoURL := os.Getenv("PROMO_SERVICE_URL")
	regularURL := os.Getenv("REGULAR_SERVICE_URL")

	var (
		promoResp   *serviceResponse
		regularResp *serviceResponse
		promoErr    error
		regularErr  error
		mu          sync.Mutex
		wg          sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		resp, err := fetchPrice(fmt.Sprintf("%s/price?product_id=%s&tier=%s", promoURL, productID, clientInfo.Tier))
		mu.Lock()
		promoResp, promoErr = resp, err
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		resp, err := fetchPrice(fmt.Sprintf("%s/price?product_id=%s&tier=%s", regularURL, productID, clientInfo.Tier))
		mu.Lock()
		regularResp, regularErr = resp, err
		mu.Unlock()
	}()

	wg.Wait()

	w.Header().Set("Content-Type", "application/json")

	if promoResp != nil && promoErr == nil {
		json.NewEncoder(w).Encode(PriceResponse{
			ProductID:  productID,
			ClientID:   clientInfo.ClientID,
			ClientName: clientInfo.Name,
			Tier:       clientInfo.Tier,
			Price:      promoResp.Price,
			Currency:   promoResp.Currency,
			Source:     "promo",
			PromoName:  promoResp.PromoName,
		})
		return
	}

	if regularResp != nil && regularErr == nil {
		json.NewEncoder(w).Encode(PriceResponse{
			ProductID:  productID,
			ClientID:   clientInfo.ClientID,
			ClientName: clientInfo.Name,
			Tier:       clientInfo.Tier,
			Price:      regularResp.Price,
			Currency:   regularResp.Currency,
			Source:     "regular",
		})
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
}

func fetchClient(url string) (*ClientInfo, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ci ClientInfo
	if err := json.Unmarshal(body, &ci); err != nil {
		return nil, err
	}

	return &ci, nil
}

func fetchPrice(url string) (*serviceResponse, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sr serviceResponse
	if err := json.Unmarshal(body, &sr); err != nil {
		return nil, err
	}

	return &sr, nil
}

func main() {
	http.HandleFunc("/get-price", getPriceHandler)
	log.Println("Price aggregator starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
