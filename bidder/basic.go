package bidder

import (
	"fmt"
	"nena-manipu-latina/models"
)

type BasicBidder struct {
	Name string
}

// TODO: Properly implement bidder interface. Implement Name()

func (b *BasicBidder) Bid(req models.BidRequest) models.BidResponse {
	// TODO: Evolve from hardcoded bids, basic logic to return a bid
	return models.BidResponse{
		AdMarkup: fmt.Sprintf("<div>Buy from %s!</div>", b.Name),
		CPM:      1.23,
		Bidder:   b.Name,
	}
}
