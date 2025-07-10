package auction

import (
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
)

func RunAuction(bidders []bidder.Bidder, strategyName string, req models.AdRequest) models.AdResponse {
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Starting for strategy [%s], request: %v", strategyName, req))

	stratImplementation, err := Factory(strategyName)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Run Auction [ERROR]: %v", err))
		return models.AdResponse{}
	}
	result, err := stratImplementation.Auction(bidders, req)
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Auction Complete. Winning Response: %v", result))
	return result
}
