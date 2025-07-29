package auction

import (
	"context"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"strings"
	"sync"
	"time"
)

/*
?---------------------------------------------------------------------------------------------------------------
 * Why this is production-grade:
 * Context-based timeout control
 * Safe writes to channel (avoids panics)
 * Partial response collection (even if some bidders fail)
 * Structured error logging per bidder (optional hook)
 * Proper channel draining
 ?---------------------------------------------------------------------------------------------------------------
*/

var collectBidsConcurrently = func(
	parentCtx context.Context,
	bidders []bidder.Bidder,
	req models.BidRequest,
	timeout time.Duration,
	perBidderTimeout time.Duration,
) models.BidCollectionResult {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()
	resCh := make(chan *models.BidResponse, len(bidders))
	errCh := make(chan error, len(bidders))

	var wg sync.WaitGroup
	latencies := make(map[string]time.Duration)
	latencyMu := sync.Mutex{}
	for _, b := range bidders {
		wg.Add(1)
		go func(b bidder.Bidder) {
			defer wg.Done()

			bidderName := b.Name()
			bidderCtx, bidderCancel := context.WithTimeout(ctx, perBidderTimeout)
			defer bidderCancel()

			start := time.Now()
			resp, err := b.Bid(req)
			latency := time.Since(start)

			select {
			case <-bidderCtx.Done():
				// ? timeout case
				errCh <- fmt.Errorf("bidder %s has timed out after %v", bidderName, perBidderTimeout)
				eight_ball_logger.Info(fmt.Sprintf("bidder %s has timed out after %v", bidderName, perBidderTimeout))
				return
			default:
				// ? check response
				if err != nil {
					errCh <- fmt.Errorf("bidder %s error: %w", bidderName, err)
					eight_ball_logger.Info(fmt.Sprintf("bidder %s error: %v", bidderName, err))
					return
				}
				latencyMu.Lock()
				latencies[bidderName] = latency
				latencyMu.Unlock()

				select {
				case resCh <- &resp:
				case <-ctx.Done():
					// ? Late response
				}
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
	case <-ctx.Done():
		// ? Time out waiting
		eight_ball_logger.Info("Concurrent Bid Collection: Timeout has been reached")
	}

	// ? Drain the channel safely
	var results []*models.BidResponse
drainResponses:
	for {
		select {
		case res := <-resCh:
			if res != nil {
				results = append(results, res)
			}
		default:
			// ? No more responses available
			break drainResponses
		}
	}

	var errors []error
drainErrors:
	for {
		select {
		case err := <-errCh:
			errors = append(errors, err)
		default:
			break drainErrors
		}
	}

	// ? Build metrics
	metrics := models.BidCollectionMetrics{
		TotalBidders:      len(bidders),
		SuccessfulBids:    len(results),
		ResponseLatencies: latencies,
	}
	for _, err := range errors {
		if strings.Contains(err.Error(), "timed out") {
			metrics.TimeoutBidders++
		} else {
			metrics.FailedBidders++
		}
	}
	return models.BidCollectionResult{
		Responses: results,
		Errors:    errors,
		Metrics:   metrics,
	}
}

// TODO:
// * integrate with Roundrobin
