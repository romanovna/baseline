package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PriceResponse struct {
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
}

var prices = map[string]map[string]float64{
	"default": {
		"1": 100.00,
		"2": 250.00,
		"3": 75.50,
		"4": 340.00,
		"5": 199.99,
	},
	"vip": {
		"1": 85.00,
		"2": 210.00,
		"3": 65.00,
		"4": 290.00,
		"5": 169.99,
	},
	"wholesale": {
		"1": 70.00,
		"2": 180.00,
		"3": 55.00,
		"4": 240.00,
		"5": 139.99,
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

	tierPrices, ok := prices[tier]
	if !ok {
		tierPrices = prices["default"]
	}

	price, ok := tierPrices[productID]
	if !ok {
		http.Error(w, `{"error":"product not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PriceResponse{
		ProductID: productID,
		Price:     price,
		Currency:  "RUB",
	})
}

func main() {
	http.HandleFunc("/price", priceHandler)
	log.Println("Regular pricing service starting on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
