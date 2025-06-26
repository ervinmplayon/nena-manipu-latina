package bidder

import (
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
	return models.AdResponse{
		Bidder:   mb.Name,
		CPM:      mb.CPM,
		AdMarkup: "<div>Mock AD from " + mb.Name + "</div>",
	}
}
