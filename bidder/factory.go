package bidder

import "time"

// TODO: replace with real logic once admin UI and dynamic config are in place
var defaultMockBidders = []Bidder{
	&MockBidder{Name: "BidderA", Delay: 50 * time.Millisecond, CPM: 1.10},
	&MockBidder{Name: "BidderB", Delay: 70 * time.Millisecond, CPM: 1.25},
	&MockBidder{Name: "BidderC", Delay: 30 * time.Millisecond, CPM: 0.95},
}
