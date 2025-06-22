package bidder

import "nena-manipu-latina/models"

type Bidder interface {
	Bid(req models.AdRequest) models.AdResponse
}
