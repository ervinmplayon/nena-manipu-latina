// ? Minumum Price Filter Bidder decorator
package bidder

import (
	"errors"
	"fmt"
	"nena-manipu-latina/models"
)

type MinPriceBidder struct {
	Inner    Bidder
	MinPrice float64
}

func (m *MinPriceBidder) Bid(req models.AdRequest) (models.AdResponse, error) {
	res, err := m.Inner.Bid(req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Min Price Bidder: Error on Bid() %s", err))
		return models.AdResponse{}, errors.New("min Price Bidder error on Bid()")
	}

	if res.CPM < m.MinPrice {
		// Return empty response if invalid or too low. NOT AN ERROR
		eight_ball_logger.Info(fmt.Sprintf("Min Price Bidder: Too cheap damn it. Min Price is %v, CPM is %v", m.MinPrice, res.CPM))
		return models.AdResponse{}, nil
	}
	return res, nil
}
