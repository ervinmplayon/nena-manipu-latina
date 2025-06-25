// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import (
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"time"
)

type FirstResponder struct{}

func (fr *FirstResponder) Auction(bids []bidder.Bidder, req models.AdRequest) models.AdResponse {
	resCh := make(chan models.AdResponse, len(bids))
	for _, bb := range bids {
		go func(bidder bidder.Bidder) {
			resCh <- bidder.Bid(req)
		}(bb)
	}
	// ? In production, I want to protect against dead bidders. If a bidder does not respond,
	// ? this can block forever, adding a timeout.
	select {
	case res := <-resCh:
		return res // * first to respond wins
	case <-time.After(10 * time.Millisecond):
		return models.AdResponse{} // * fallback / timeout
	}
}
