package bidder

import "nena-manipu-latina/models"

type Bidder interface {
	Name() string
	Bid(req models.BidRequest) (models.BidResponse, error)
}
