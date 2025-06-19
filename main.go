package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nena-manipu-latina/models"
	"net/http"
	"time"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

// ! Because 8-balls are neither good nor bad, its just mid asf
var eight_ball_logger logger.Logger = &logger.LogrusLogger{}

func handleAdRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != http.MethodPost {
		/*
		 * We are literally learning on the fly here
		 * TODO: (move me somewhere like in a bible and give me dignity, wicked sorcerer)
		 * IT is a standard paractive to log errors like malfromed requests.
		 * `http.Error` writes the response to the client - but it doesn't tell now us (the server)
		 * now does it? This is where the logger comes in.
		 */
		eight_ball_logger.Error(fmt.Sprintf("Invalid method: %s", r.Method))
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
