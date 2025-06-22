// ? Selects the winner based on round-robin. Usefule for equal distribution
package auction

import (
	"nena-manipu-latina/models"
	"sync"
)

type RoundRobin struct {
	mu    sync.Mutex
	index int
}

func (rr *RoundRobin) Auction(bids []models.AdResponse) models.AdResponse {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if len(bids) == 0 {
		return models.AdResponse{}
	}

	selected := bids[rr.index%len(bids)]
	rr.index++
	return selected
}
