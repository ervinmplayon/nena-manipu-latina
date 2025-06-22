package bidder

import (
	"fmt"
	"nena-manipu-latina/models"
)

type BasicBidder struct {
	Name string
}

func (b *BasicBidder) Bid(req models.AdRequest) models.AdResponse {
	// TODO: Evolve from hardcoded bids, basic logic to return a bid
	return models.AdResponse{
		AdMarkup: fmt.Sprintf("<div>Buy from %s!</div>", b.Name),
		CPM:      1.23,
		Bidder:   b.Name,
	}
}
