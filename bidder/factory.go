package bidder

import (
	"time"
)

// TODO: replace with real logic once admin UI and dynamic config are in place
var defaultMockBidders = []Bidder{
	NewMockBidder("BidderA", 50*time.Millisecond, 1.10),
	NewMockBidder("BidderB", 70*time.Millisecond, 1.25),
	NewMockBidder("BidderC", 30*time.Millisecond, 0.95),
}

/*
 * Factory returns a list of bidders based on a publisher ID.
 * For now, it's hardcoded, later query a DB or config service.
 */
var Factory = func(publisherID string) []Bidder {
	switch publisherID {
	case "publisher_firstresponder":
		eight_ball_logger.Info("Bidder Strategy Factory: Selected FirstResponder Bidders")
		return []Bidder{
			NewMockBidder("FastDSP", 20*time.Millisecond, 1.05),
			NewMockBidder("ReliableDSP", 40*time.Millisecond, 1.10),
		}
	case "publisher_roundrobin":
		eight_ball_logger.Info("Bidder Strategy Factory: Selected RoundRobin Bidders")
		return []Bidder{
			NewMockBidder("SlowAssBidder", 100*time.Millisecond, 0.99),
			NewMockBidder("AggressiveDSP", 50*time.Millisecond, 1.30),
		}
	default:
		// ? Use defaultMockBidders as fallback
		eight_ball_logger.Info("Bidder Strategy Factory: Selected DEFAULT. Using Mock Bidders")
		return defaultMockBidders
	}
}
