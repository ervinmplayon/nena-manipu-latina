package auction

import "nena-manipu-latina/models"

type AuctionStrategy interface {
	// ? this will make sense as an auction takes ad responses and have 1 ad win
	Auction([]models.AdResponse) models.AdResponse
}
