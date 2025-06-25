// ? Selects the winner based on round-robin. Usefule for equal distribution
package auction

import (
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"sync/atomic"
)

type RoundRobin struct {
	index atomic.Uint64
}

func (rr *RoundRobin) Auction(bids []bidder.Bidder, req models.AdRequest) models.AdResponse {
	if len(bids) == 0 {
		return models.AdResponse{}
	}
	i := rr.index.Add(1)
	selected := bids[i%uint64(len(bids))]
	return selected
}
