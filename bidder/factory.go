package bidder

import (
	"time"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

// TODO: replace with real logic once admin UI and dynamic config are in place
var defaultMockBidders = []Bidder{
	&MockBidder{Name: "BidderA", Delay: 50 * time.Millisecond, CPM: 1.10},
	&MockBidder{Name: "BidderB", Delay: 70 * time.Millisecond, CPM: 1.25},
	&MockBidder{Name: "BidderC", Delay: 30 * time.Millisecond, CPM: 0.95},
}

// ? This is declared at the package level, it has package scope. It can be used anywhere in this package
var eight_ball_logger logger.Logger = &logger.StandardLogger{}

/*
 * Factory returns a list of bidders based on a publisher ID.
 * For now, it's hardcoded, later query a DB or config service.
 */
func Factory(publisherID string) []Bidder {
	switch publisherID {
	case "publisher_firstresponder":
		eight_ball_logger.Info("Bidder Strategy Factory: Selected FirstResponder")
		return []Bidder{
			&MockBidder{Name: "FastDSP", Delay: 20 * time.Millisecond, CPM: 1.05},
			&MockBidder{Name: "ReliableDSP", Delay: 40 * time.Millisecond, CPM: 1.10},
		}
	case "publisher_roundrobin":
		eight_ball_logger.Info("Bidder Strategy Factory: Selected RoundRobin")
		return []Bidder{
			&MockBidder{Name: "SlowBidder", Delay: 100 * time.Millisecond, CPM: 0.99},
			&MockBidder{Name: "AggressiveDSP", Delay: 50 * time.Millisecond, CPM: 1.30},
		}
	default:
		// ? Use defaultMockBidders as fallback
		eight_ball_logger.Info("Bidder Strategy Factory: Selected DEFAULT. Using Mock Bidders")
		return defaultMockBidders
	}
}
