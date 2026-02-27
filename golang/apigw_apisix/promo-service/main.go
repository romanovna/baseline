package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PromoResponse struct {
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	PromoName string  `json:"promo_name"`
}

type promoEntry struct {
	Price     float64
	PromoName string
}

var promos = map[string]map[string]promoEntry{
	"default": {
		"1": {Price: 79.99, PromoName: "Summer Sale"},
		"3": {Price: 49.99, PromoName: "Flash Deal"},
	},
	"vip": {
		"1": {Price: 69.99, PromoName: "VIP Summer"},
		"3": {Price: 39.99, PromoName: "VIP Flash"},
		"5": {Price: 149.99, PromoName: "VIP Exclusive"},
	},
}

func priceHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	if productID == "" {
		http.Error(w, `{"error":"product_id is required"}`, http.StatusBadRequest)
		return
	}

	tier := r.URL.Query().Get("tier")
	if tier == "" {
		tier = "default"
	}

	tierPromos, ok := promos[tier]
	if !ok {
		tierPromos = promos["default"]
	}

	entry, ok := tierPromos[productID]
	if !ok {
		http.Error(w, `{"error":"no promo for this product"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PromoResponse{
		ProductID: productID,
		Price:     entry.Price,
		Currency:  "RUB",
		PromoName: entry.PromoName,
	})
}

func main() {
	http.HandleFunc("/price", priceHandler)
	log.Println("Promo pricing service starting on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
