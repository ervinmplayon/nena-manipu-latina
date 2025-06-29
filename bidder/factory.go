package bidder

import "time"

// TODO: replace with real logic once admin UI and dynamic config are in place
var defaultMockBidders = []Bidder{
	&MockBidder{Name: "BidderA", Delay: 50 * time.Millisecond, CPM: 1.10},
	&MockBidder{Name: "BidderB", Delay: 70 * time.Millisecond, CPM: 1.25},
	&MockBidder{Name: "BidderC", Delay: 30 * time.Millisecond, CPM: 0.95},
}

/*
 * Factory returns a list of bidders based on a publisher ID.
 * For now, it's hardcoded, later query a DB or config service.
 */
func Factory(publisherID string) []Bidder {
	switch publisherID {
	case "publisher_firstresponder":
		return []Bidder{
			&MockBidder{Name: "FastDSP", Delay: 20 * time.Millisecond, CPM: 1.05},
			&MockBidder{Name: "ReliableDSP", Delay: 40 * time.Millisecond, CPM: 1.10},
		}
	case "publisher_roundrobin":
		return []Bidder{
			&MockBidder{Name: "SlowBidder", Delay: 100 * time.Millisecond, CPM: 0.99},
			&MockBidder{Name: "AggressiveDSP", Delay: 50 * time.Millisecond, CPM: 1.30},
		}
	default:
		// ? Use defaultMockBidders as fallback
		return defaultMockBidders
	}
}
