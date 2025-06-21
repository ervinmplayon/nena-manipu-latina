// ? Minumum Price Filter Bidder decorator
package bidder

import "nena-manipu-latina/models"

type MinPriceBidder struct {
	Inner    Bidder
	MinPrice float64
}

func (m *MinPriceBidder) Bid(req models.AdRequest) models.AdResponse {
	res := m.Inner.Bid(req)
	if res.CPM < m.MinPrice {
		return models.AdResponse{} // empty = no bid
	}
	return res
}
