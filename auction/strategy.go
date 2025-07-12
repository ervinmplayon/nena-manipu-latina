package auction

import (
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
)

type AuctionStrategy interface {
	// ? this will make sense as an auction takes ad responses and have 1 ad win
	Auction(bids []bidder.Bidder, req models.BidRequest) (models.BidResponse, error)
}
