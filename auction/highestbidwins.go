package auction

import (
	"nena-manipu-latina/models"
)

type HighestBidWins struct{}

func (h *HighestBidWins) Auction(bids []models.BidResponse) models.BidResponse {
	var winner models.BidResponse
	var maxCPM float64

	for _, bid := range bids {
		if bid.CPM > maxCPM {
			maxCPM = bid.CPM
			winner = bid
		}
	}
	return winner
}
