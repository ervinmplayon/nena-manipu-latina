// ? Minumum Price Filter Bidder decorator
package bidder

import (
	"fmt"
	"nena-manipu-latina/models"
)

type MinPriceBidder struct {
	Inner    Bidder
	MinPrice float64
}

func (m *MinPriceBidder) Bid(req models.AdRequest) models.AdResponse {
	res, err := m.Inner.Bid(req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Min Price Bidder: Error on Bid() %s", err))
		return models.AdResponse{}
	}

	if res.CPM < m.MinPrice {
		// Return empty response if invalid or too low
		return models.AdResponse{}
	}
	return res
}
