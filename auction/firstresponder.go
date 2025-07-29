// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import (
	"context"
	"errors"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"time"
)

type FirstResponder struct {
	Timeout          time.Duration
	PerBidderTimeout time.Duration
}

func NewFirstResponder(timeout, perBidderTimeout time.Duration) *FirstResponder {
	return &FirstResponder{
		Timeout:          timeout,
		PerBidderTimeout: perBidderTimeout,
	}
}

func (f *FirstResponder) AuctionWithContext(
	ctx context.Context,
	bidders []bidder.Bidder,
	req models.BidRequest,
) (*models.BidCollectionResult, error) {
	result := collectBidsConcurrently(ctx, bidders, req, f.Timeout, f.PerBidderTimeout)

	for _, bid := range result.Responses {
		if bid != nil && bid.Bidder != "" {
			// ? First successful bidder found
			return &models.BidCollectionResult{
				Responses: []*models.BidResponse{bid},
				Errors:    result.Errors,
				Metrics:   result.Metrics,
			}, nil
		}
	}

	eight_ball_logger.Error("first Responder: no bids received")
	return nil, errors.New("first Responder: no bids received")
}
