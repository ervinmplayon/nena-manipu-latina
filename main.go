package main

import (
	"encoding/json"
	"log"
	"nena-manipu-latina/models"
	"net/http"
	"time"
)

func handleAdRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var adReq models.AdRequest
	if err := json.NewDecoder(r.Body).Decode(&adReq); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Received request: %v\n", adReq)

	// TODO: Auction, creative, logging
	adResp := models.AdResponse{
		AdMarkup: "<div>Your AD Here</div>",
		CPM:      "1.25",
		Bidder:   "mock-bidder",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adResp)

	log.Printf("Responded in %v\n", time.Since(start))
}

func main() {
	http.HandleFunc("/serve", handleAdRequest)
	port := ":8080"
	log.Printf("Starting ad server on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
