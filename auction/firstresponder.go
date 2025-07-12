// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import (
	"errors"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"time"
)

type FirstResponder struct{}

func (fr *FirstResponder) Auction(bids []bidder.Bidder, req models.BidRequest) (models.BidResponse, error) {
	if len(bids) == 0 {
		return models.BidResponse{}, errors.New("first Responder: no bidders available")
	}

	type result struct {
		resp models.BidResponse
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
				eight_ball_logger.Info(fmt.Sprintf("First Responder: %s won the bid", res.resp.Bidder))
				return res.resp, nil // * first to respond wins
			}
		case <-timeout:
			return models.BidResponse{}, errors.New("first Responder: timeout waiting for bids") // * fallback / timeout
		}
	}
	return models.BidResponse{}, errors.New("first Responder: no valid bids returned")
	// TODO: rethink if returning an error makes the most sense. Perhaps an Info will suffice?
	// * Apply the same to `case<-timeout`
}
