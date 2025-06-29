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
	// * Placeholder: customize based on publisherID later
	return defaultMockBidders
}
