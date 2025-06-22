package auction

import (
	"nena-manipu-latina/models"
)

type HighestBidWins struct{}

func (h *HighestBidWins) Auction(bids []models.AdResponse) models.AdResponse {
	var winner models.AdResponse
	var maxCPM float64

	for _, bid := range bids {
		if bid.CPM > maxCPM {
			maxCPM = bid.CPM
			winner = bid
		}
	}
	return winner
}
