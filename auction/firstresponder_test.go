package auction_test

import (
	"context"
	"errors"
	"nena-manipu-latina/models"
	"testing"
	"time"
)

// ? Implement the bidder.Bidder interface as mock
type mockBidder struct {
	name      string
	shouldBid bool
	delay     time.Duration
}

func (m *mockBidder) Name() string {
	return m.name
}

func (m *mockBidder) Bid(ctx context.Context, req models.BidRequest) (*models.BidResponse, error) {
	select {
	case <-time.After(m.delay):
		if m.shouldBid {
			return &models.BidResponse{
				Creative: "creative-" + m.name,
				CPM:      1.23,
				Bidder:   m.name,
			}, nil
		}
		return nil, errors.New("no bid")
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func TestFirstResponder_AuctionWithContext(t *testing.T) {
	// TODO
	// * Arrange
	// bidders := []bidder.Bidder{
	// 	&mockBidder{name: "slow1", delay: 300 * time.Millisecond, shouldBid: true},
	// 	&mockBidder{name: "fast1", delay: 50 * time.Millisecond, shouldBid: true}, // <- should win
	// 	&mockBidder{name: "slow2", delay: 400 * time.Millisecond, shouldBid: true},
	// 	&mockBidder{name: "nobid", delay: 50 * time.Millisecond, shouldBid: false}, // <- no bid
	// }
	// * Act

	// * Assert

}
