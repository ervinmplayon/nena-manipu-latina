package auction_test

import (
	"context"
	"errors"
	"nena-manipu-latina/auction"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	// * Arrange
	bidders := []bidder.Bidder{
		&mockBidder{name: "slow1", delay: 300 * time.Millisecond, shouldBid: true},
		&mockBidder{name: "fast1", delay: 50 * time.Millisecond, shouldBid: true}, // <- should win
		&mockBidder{name: "slow2", delay: 400 * time.Millisecond, shouldBid: true},
		&mockBidder{name: "nobid", delay: 50 * time.Millisecond, shouldBid: false}, // <- no bid
	}
	req := models.BidRequest{
		RequestID:   "test-auction",
		PublisherID: "test-auction",
	}

	// * Set the parent timeout to be longer than per-bidder timeout
	ctx, cancel := context.WithTimeout(t.Context(), 150*time.Millisecond)
	defer cancel()

	strategy := &auction.FirstResponder{}

	// * Act
	result, err := strategy.AuctionWithContext(ctx, bidders, req)

	// * Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Len(t, result.Responses, 1)
	assert.Equal(t, "fast1", result.Responses[0].Bidder)
	assert.Equal(t, "creative-fast1", result.Responses[0].Creative)

	t.Logf("Winner: %+v", result.Responses[0])
	t.Logf("Errors: %+v", result.Errors)
}
