package bidder

import (
	"context"
	"nena-manipu-latina/models"
)

type Bidder interface {
	Name() string
	Bid(ctx context.Context, req models.BidRequest) (models.BidResponse, error)
}
