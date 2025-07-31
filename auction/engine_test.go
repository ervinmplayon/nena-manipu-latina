package auction_test

import (
	"context"
	"nena-manipu-latina/auction"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
	"testing"
	"time"
)

// ? Define mock strategy for testing
type mockStrategy struct{}

func (m *mockStrategy) AuctionWithContext(ctx context.Context, bidders []bidder.Bidder, req models.BidRequest) (*models.BidCollectionResult, error) {
	mockResponse := models.BidResponse{
		Bidder: "TestDSP",
		CPM:    1.23,
	}
	return &models.BidCollectionResult{
		Responses: []*models.BidResponse{&mockResponse},
		Errors:    nil,
		Metrics:   models.BidCollectionMetrics{},
	}, nil
}

// ? Override the Factory function for testing
func mockFactory(strategyName string) (auction.AuctionStrategyWithContext, error) {
	return &mockStrategy{}, nil
}

func TestRunAuction_ReturnsExpectedResponse(t *testing.T) {
	// * temporarily swap out the Factory
	originalFactory := auction.Factory
	auction.Factory = mockFactory
	defer func() { auction.Factory = originalFactory }() // ? Restore after the test

	// ? Create mock bidders (though this mockStrategy ignores them)
	bidders := []bidder.Bidder{
		bidder.NewMockBidder("Mock1", 10*time.Millisecond, 1.00),
	}

	// * Call RunAuction
	req := models.BidRequest{AdUnit: "banner"}
	resp, _ := auction.RunAuction(context.Background(), bidders, "any-strategy", req)

	if resp.Responses[0].Bidder != "TestDSP" || resp.Responses[0].CPM != 1.23 {
		t.Errorf("Unexpected auction result: %+v", resp)
	}
}
