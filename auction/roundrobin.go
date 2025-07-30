// ? Selects the winner based on round-robin. Usefule for equal distribution
package auction

import (
	"context"
	"errors"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"sync/atomic"
)

type RoundRobin struct {
	index atomic.Uint64
}

func (rr *RoundRobin) AuctionWithContext(ctx context.Context, bids []bidder.Bidder, req models.BidRequest) (*models.BidResponse, error) {
	if len(bids) == 0 {
		return &models.BidResponse{}, errors.New("round Robin: no bidders available")
	}
	i := rr.index.Add(1)
	selected := bids[i%uint64(len(bids))]
	resp, err := selected.Bid(ctx, req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("bidder %s failed: %s", selected.Name(), err.Error()))
		return &models.BidResponse{}, errors.New("round Robin bidder error")
	}
	eight_ball_logger.Info(fmt.Sprintf("Round Robin: %s won the bid", resp.Bidder))
	return resp, nil
}
