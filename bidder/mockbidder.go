package bidder

import (
	"fmt"
	"nena-manipu-latina/models"
	"time"
)

type MockBidder struct {
	Name  string
	Delay time.Duration
	CPM   float64
}

func (mb *MockBidder) Bid(req models.AdRequest) models.AdResponse {
	time.Sleep(mb.Delay)
	eight_ball_logger.Info(fmt.Sprintf("[MOCK BIDDER] %s responding with CPM=%.2f after %s", mb.Name, mb.CPM, mb.Delay))
	return models.AdResponse{
		Bidder:   mb.Name,
		CPM:      mb.CPM,
		AdMarkup: "<div>Mock AD from " + mb.Name + "</div>",
	}
}
