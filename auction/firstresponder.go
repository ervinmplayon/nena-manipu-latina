// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import (
	"errors"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"time"
)

type FirstResponder struct{}

func (fr *FirstResponder) Auction(bids []bidder.Bidder, req models.AdRequest) (models.AdResponse, error) {
	if len(bids) == 0 {
		return models.AdResponse{}, errors.New("first Responder: no bidders available")
	}

	type result struct {
		resp models.AdResponse
		err  error
	}

	resCh := make(chan result, len(bids))
	for _, bb := range bids {
		go func(bb bidder.Bidder) {
			resp, err := bb.Bid(req)
			resCh <- result{resp, err}
		}(bb)
	}
	// ? In production, I want to protect against dead bidders. If a bidder does not respond,
	// ? this can block forever, adding a timeout aka Defensive Handling
	timeout := time.After(50 * time.Millisecond)
	for range bids {
		select {
		case res := <-resCh:
			if res.err == nil {
				return res.resp, nil // * first to respond wins
			}
		case <-timeout:
			return models.AdResponse{}, errors.New("first Responder: timeout waiting for bids") // * fallback / timeout
		}
	}
	return models.AdResponse{}, errors.New("first Responder: no valid bids returned")
}
