// ? Minumum Price Filter Bidder decorator
package bidder

import (
	"nena-manipu-latina/models"
	"strconv"
)

type MinPriceBidder struct {
	Inner    Bidder
	MinPrice float64
}

func (m *MinPriceBidder) Bid(req models.AdRequest) models.AdResponse {
	res := m.Inner.Bid(req)
	cpm, err := strconv.ParseFloat(res.CPM, 64)
	if err != nil || cpm < m.MinPrice {
		// Return empty response if invalid or too low
		return models.AdResponse{}
	}
	return res
}
