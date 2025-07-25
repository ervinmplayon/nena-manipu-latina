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

func (fr *FirstResponder) Auction(bidders []bidder.Bidder, req models.BidRequest) (models.BidResponse, error) {
	results := collectBidsConcurrently(bidders, req, 50*time.Millisecond)

	for _, res := range results {
		if res.Err == nil && res.Response != nil {
			eight_ball_logger.Info(fmt.Sprintf("First Responder: %s won the bid", res.Response.Bidder))
			return *res.Response, nil // * First valid response
		}
	}
	eight_ball_logger.Error("first responder: No valid bidder response")
	return models.BidResponse{}, errors.New("first responder: No valid bidder response")
}
