// ? Maximum Price Filter Bdder decorator
package bidder

import (
	"fmt"
	"nena-manipu-latina/models"
)

type MaxPriceBidder struct {
	Inner    Bidder
	MaxPrice float64
}

func (m *MaxPriceBidder) Bid(req models.AdRequest) (models.AdResponse, error) {
	res, err := m.Inner.Bid(req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Max Price Bidder: Error on Bid() %s", err))
		return models.AdResponse{}, err
	}

	if res.CPM > m.MaxPrice {
		// Return empty response if invalid or too high. NOT AN ERROR
		eight_ball_logger.Info(fmt.Sprintf("Max Price Bidder: Too damn high. Max Price is %v, CPM is %v", m.MaxPrice, res.CPM))
		return models.AdResponse{}, nil
	}
	return res, nil
}
