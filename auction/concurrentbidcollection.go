package auction

import (
	"context"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"sync"
	"time"
)

/*
 * Why this is production-grade:
 * Context-based timeout control
 * Safe writes to channel (avoids panics)
 * Partial response collection (even if some bidders fail)
 * Structured error logging per bidder (optional hook)
 */

var collectBidsConcurrently = func(
	ctx context.Context,
	bidders []bidder.Bidder,
	req models.BidRequest,
	timeout time.Duration,
) []*models.BidResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	resCh := make(chan *models.BidResponse, len(bidders))

	var wg sync.WaitGroup
	for _, b := range bidders {
		wg.Add(1)
		go func(b bidder.Bidder) {
			defer wg.Done()
			// ? Optional: wrap the bidder call with per-bidder timeout if needed
			resp, err := b.Bid(req)
			if err != nil {
				eight_ball_logger.Info(fmt.Sprintf("Concurrent Bid Collection: Bidder %s error: %v", b.Name(), err))
				return
			}
			// ? Safe send with select - avoids panics is ctx is done
			select {
			case resCh <- &resp:
			case <-ctx.Done():
				// ? Too late to send, main process is aborting
				eight_ball_logger.Info("Concurrent Bid Collection: <-ctx.Done() has been reached")
			}
		}(b)
	}

	// ? Wait for all bidders to finish in a separate goroutine
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// ? Block until either all responses are in, or a timeout hits
	select {
	// ? `<-done` is instantly successful once `close(done)` is called.
	case <-done:
		// ? All responses collected
		eight_ball_logger.Info("Concurrent Bid Collection: All bidders completed")
	case <-time.After(timeout):
		// ? Time out waiting
		eight_ball_logger.Info("Concurrent Bid Collection: Timeout has been reached")
	}

	// ? Drain the channel safely
	var results []*models.BidResponse
	for {
		select {
		case res := <-resCh:
			if res != nil {
				results = append(results, res)
			}
		default:
			// ? No more responses available
			return results
		}
	}
}
