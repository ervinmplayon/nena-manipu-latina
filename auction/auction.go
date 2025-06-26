package auction

import (
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
)

func RunAuction(bidders []bidder.Bidder, req models.AdRequest) models.AdResponse {
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Auction starting for request: %v", req))
}
