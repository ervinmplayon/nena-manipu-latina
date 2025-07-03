package auction_test

import (
	"nena-manipu-latina/auction"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"testing"
	"time"
)

// ? Define mock strategy for testing
type mockStrategy struct{}

func (m *mockStrategy) Auction(bidders []bidder.Bidder, req models.AdRequest) models.AdResponse {
	return models.AdResponse{
		Bidder: "TestDSP",
		CPM:    1.23,
	}
}

// ? Override the Factory function for testing
func mockFactory(strategyName string) (auction.AuctionStrategy, error) {
	return &mockStrategy{}, nil
}

func TestRunAuction_ReturnsExpectedResponse(t *testing.T) {
	// * temporarily swap out the Factory
	originalFactory := auction.Factory
	auction.Factory = mockFactory
	defer func() { auction.Factory = originalFactory }() // ? Restore after the test

	// ? Create mock bidders (though this mockStrategy ignores them)
	bidders := []bidder.Bidder{
		&bidder.MockBidder{Name: "Mock1", Delay: 10 * time.Millisecond, CPM: 1.00},
	}

	// * Call RunAuction
	req := models.AdRequest{AdUnit: "banner"}
	resp := auction.RunAuction(bidders, "any-strategy", req)

	if resp.Bidder != "TestDSP" || resp.CPM != 1.23 {
		t.Errorf("Unexpected auction result: %+v", resp)
	}
}
