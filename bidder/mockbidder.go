package bidder

import (
	"fmt"
	"nena-manipu-latina/models"
	"time"
)

type MockBidder struct {
	name  string
	Delay time.Duration
	CPM   float64
}

func NewMockBidder(name string, delay time.Duration, cpm float64) *MockBidder {
	return &MockBidder{
		name:  name,
		Delay: delay,
		CPM:   cpm,
	}
}

func (mb *MockBidder) Name() string {
	return mb.name
}

func (mb *MockBidder) Bid(req models.BidRequest) (models.BidResponse, error) {
	time.Sleep(mb.Delay)
	eight_ball_logger.Info(fmt.Sprintf("[MOCK BIDDER] %s responding with CPM=%.2f after %s", mb.name, mb.CPM, mb.Delay))
	return models.BidResponse{
		Bidder:   mb.name,
		CPM:      mb.CPM,
		Creative: "<div>Mock AD from " + mb.name + "</div>",
	}, nil
}
