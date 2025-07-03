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
		mockBidder{
			name:          "BidderA",
			response:      models.AdResponse{CPM: 2.50, AdMarkup: "A content", Bidder: "A"},
			responseDelay: 30 * time.Millisecond,
		},
		mockBidder{
			name:          "BidderB",
			response:      models.AdResponse{CPM: 1.00, AdMarkup: "B content", Bidder: "B"},
			responseDelay: 10 * time.Millisecond, // ? <-- will win
		},
		mockBidder{
			name:          "BidderC",
			response:      models.AdResponse{CPM: 5.00, AdMarkup: "C content", Bidder: "C"},
			responseDelay: 50 * time.Millisecond,
		},
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
	if resp.AdMarkup != "B content" || resp.Bidder != "B" {
		t.Errorf("Expected B content from BidderB, got %+v", resp)
	}
}

// ? Test when no bidders respond
func TestFirstResponder_NoValidBids(t *testing.T) {
	strategy := auction.FirstResponder{}

	bidders := []bidder.Bidder{
		mockBidder{name: "BidderX", response: models.AdResponse{}},
		mockBidder{name: "BidderY", response: models.AdResponse{}},
	}

	req := models.AdRequest{
		RequestID:   "test456",
		PublisherID: "pub2",
		AdUnit:      "unit2",
		DeviceIP:    "5.6.7.8",
	}

	resp := strategy.Auction(bidders, req)

	if resp != (models.AdResponse{}) {
		t.Errorf("Expected empty response, got %+v", resp)
	}
}
