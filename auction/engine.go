package auction

import (
	"errors"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
)

func RunAuction(bidders []bidder.Bidder, strategyName string, req models.BidRequest) (models.BidResponse, error) {
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Starting for strategy [%s], request: %v", strategyName, req))

	stratImplementation, err := Factory(strategyName)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Run Auction [ERROR] after Factory: %v", err))
		return models.BidResponse{}, errors.New("run Auction error after Factory")
	}
	result, err := stratImplementation.Auction(bidders, req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Run Auction [ERROR] after Auction: %v", err))
		return models.BidResponse{}, errors.New("run Auction error after Auction")
	}
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Auction Complete. Winning Response: %v", result))
	return result, nil
}
