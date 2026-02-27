package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
	Tier     string `json:"tier"`
}

var clients = map[string]Client{
	"client_001": {ClientID: "client_001", Name: "ООО Ромашка", Tier: "default"},
	"client_002": {ClientID: "client_002", Name: "ИП Иванов", Tier: "vip"},
	"client_003": {ClientID: "client_003", Name: "ООО ОптТорг", Tier: "wholesale"},
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		http.Error(w, `{"error":"client_id is required"}`, http.StatusBadRequest)
		return
	}

	client, ok := clients[clientID]
	if !ok {
		http.Error(w, `{"error":"client not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func main() {
	http.HandleFunc("/client", clientHandler)
	log.Println("Client service starting on :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
