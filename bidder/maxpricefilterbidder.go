// ? Maximum Price Filter Bdder decorator
package bidder

import (
	"nena-manipu-latina/models"
)

type MaxPriceBidder struct {
	Inner    Bidder
	MaxPrice float64
}

func (m *MaxPriceBidder) Bid(req models.AdRequest) models.AdResponse {
	res := m.Inner.Bid(req)

	if res.CPM > m.MaxPrice {
		// Return empty response if invalid or too high
		return models.AdResponse{}
	}
	return res
}
