// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import (
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
)

type FirstResponder struct{}

func (fr *FirstResponder) Auction(bids []bidder.Bidder, req models.AdRequest) models.AdResponse {
	resCh := make(chan models.AdResponse, len(bids))
	for _, bb := range bids {
		go func(bidder bidder.Bidder) {
			resCh <- bidder.Bid(req)
		}(bb)
	}
	return <-resCh
}
