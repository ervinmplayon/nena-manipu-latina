package auction_test

import (
	"nena-manipu-latina/auction"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"testing"
	"time"
)

// ? Implement the bidder.Bidder interface as mock
type mockBidder struct {
	name          string
	response      models.AdResponse
	responseDelay time.Duration
}

func (m mockBidder) Name() string {
	return m.name
}

func (m mockBidder) Bid(req models.AdRequest) models.AdResponse {
	if m.responseDelay > 0 {
		time.Sleep(m.responseDelay)
	}
	return m.response
}

func TestFirstResponder_Auction(t *testing.T) {
	// * Arrange
	strategy := auction.FirstResponder{}

	bidders := []bidder.Bidder{
		mockBidder{name: "BidderA", response: models.AdResponse{}}, // empty bid
		mockBidder{name: "BidderB", response: models.AdResponse{CPM: 2.50, AdMarkup: "B content"}},
		mockBidder{name: "BidderC", response: models.AdResponse{CPM: 1.00, AdMarkup: "C content"}},
	}

	req := models.AdRequest{
		RequestID:   "susie-stellar-request-123",
		PublisherID: "gina-valentina",
		AdUnit:      "sinatra-monroe-cheeks",
		DeviceIP:    "1.2.3.4",
	}

	// * Act
	resp := strategy.Auction(bidders, req)

	// * Assert
	if resp.CPM != 2.50 || resp.AdMarkup != "B content" {
		t.Errorf("Expected B content with 2.50 bid, got %+v", resp)
	}
}
