// ? Selects the winner based on round-robin. Usefule for equal distribution
package auction

import (
	"errors"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"sync/atomic"
)

type RoundRobin struct {
	index atomic.Uint64
}

func (rr *RoundRobin) Auction(bids []bidder.Bidder, req models.AdRequest) (models.AdResponse, error) {
	if len(bids) == 0 {
		return models.AdResponse{}, errors.New("round Robin: no bidders available")
	}
	i := rr.index.Add(1)
	selected := bids[i%uint64(len(bids))]
	resp, err := selected.Bid(req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("bidder %s failed: %s", selected.Name(), err.Error()))
		return models.AdResponse{}, errors.New("round Robin bidder error")
	}
	return resp, nil
}
