package auction

import (
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"sync"
	"time"
)

var collectBidsConcurrently = func(bidders []bidder.Bidder, req models.BidRequest, timeout time.Duration) []*models.AuctionResult {
	var wg sync.WaitGroup
	resCh := make(chan *models.AuctionResult, len(bidders))

	for _, b := range bidders {
		wg.Add(1)
		go func(b bidder.Bidder) {
			defer wg.Done()
			resp, err := b.Bid(req)
			resCh <- &models.AuctionResult{
				BidderName: b.Name(),
				Response:   &resp,
				Err:        err,
			}
		}(b)
	}

	// ? Wait with timeout to avoid deadlock
	// TODO: learn the internals of this whole goroutine
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	//TODO: ============================================

	var results []*models.AuctionResult
	select {
	case <-done:
		// ? All responses collected
	case <-time.After(timeout):
		// ? Time out waiting
	}

	// ? Drain the channel
	for i := 0; i < len(bidders); i++ {
		select {
		case r := <-resCh:
			results = append(results, r)
		default:
		}
	}
	return results
}
