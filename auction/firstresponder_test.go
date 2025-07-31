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

type mockBidder struct {
	name       string
	delay      time.Duration
	resp       *models.BidResponse
	err        error
	shouldHang bool
}

func (m *mockBidder) Name() string {
	return m.name
}

func (m *mockBidder) Bid(ctx context.Context, req models.BidRequest) (*models.BidResponse, error) {
	if m.shouldHang {
		<-ctx.Done()
		return nil, ctx.Err()
	}

	select {
	case <-time.After(m.delay):
		return m.resp, m.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func TestFirstResponder_AuctionWithContext(t *testing.T) {
	// * Arrange
	bidders := []bidder.Bidder{
		&mockBidder{
			name:  "fast_success",
			delay: 10 * time.Millisecond,
			resp: &models.BidResponse{
				Bidder: "fast_success",
			},
		},
		&mockBidder{
			name:  "slow_success",
			delay: 200 * time.Millisecond,
			resp: &models.BidResponse{
				Bidder: "slow_success",
			},
		},
		&mockBidder{
			name:  "error_bidder",
			delay: 50 * time.Millisecond,
			err:   errors.New("some error"),
		},
		&mockBidder{
			name:       "hanging_bidder",
			shouldHang: true,
		},
	}

	timeout := 300 * time.Millisecond
	perBidderTimeout := 100 * time.Millisecond
	strategy := auction.NewFirstResponder(timeout, perBidderTimeout)

	ctx := context.Background()
	req := models.BidRequest{
		RequestID:   "test-auction",
		PublisherID: "test-auction",
	}

	result, err := strategy.AuctionWithContext(ctx, bidders, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Responses, 1)
	assert.Equal(t, "fast_success", result.Responses[0].Bidder)

	assert.True(t, result.Metrics.TotalBidders == 4)
	assert.True(t, result.Metrics.SuccessfulBids == 1)
	assert.True(t, result.Metrics.TimeoutBidders >= 1) // hanging bidder
	assert.True(t, result.Metrics.FailedBidders >= 1)  // error_bidder
}
