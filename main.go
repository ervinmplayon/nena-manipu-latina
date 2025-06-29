package main

import (
	"encoding/json"
	"fmt"
	"nena-manipu-latina/auction"
	"nena-manipu-latina/config"
	"nena-manipu-latina/models"
	"net/http"
	"time"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

// ? Because 8-balls are neither good nor bad, its just mid asf
var eight_ball_logger logger.Logger = &logger.LogrusLogger{}

func handleAdRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != http.MethodPost {
		eight_ball_logger.Error(fmt.Sprintf("Invalid method: %s", r.Method))
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var adReq models.AdRequest
	if err := json.NewDecoder(r.Body).Decode(&adReq); err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Failed to decode JSON: %s", err.Error()))
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	eight_ball_logger.Info(fmt.Sprintf("Received Request from Publisher: %s\n", adReq.PublisherID))
	pubConfig, found := config.GetPublisherConfig(adReq.PublisherID)
	if !found {
		eight_ball_logger.Error(fmt.Sprintf("Unknown publisher ID: %s", adReq.PublisherID))
		http.Error(w, "Unknown publisher", http.StatusBadRequest)
		return
	}

	// ? Auction entry point
	adResp := auction.RunAuction(pubConfig.Bidders, pubConfig.Strategy, adReq)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adResp)

	eight_ball_logger.Info(fmt.Sprintf("Responded in %v\n", time.Since(start)))
}

func main() {
	http.HandleFunc("/serve", handleAdRequest)
	port := ":8080"
	eight_ball_logger.Info(fmt.Sprintf("Starting ad server on %s...\n", port))
	eight_ball_logger.Fatal(http.ListenAndServe(port, nil).Error())
}
