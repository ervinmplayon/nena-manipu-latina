package bidder

import "nena-manipu-latina/models"

type Bidder interface {
	Name() string
	Bid(req models.AdRequest) (models.AdResponse, error)
}
