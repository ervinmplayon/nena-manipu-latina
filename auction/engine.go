package auction

import (
	"context"
	"errors"
	"fmt"
	"nena-manipu-latina/bidder"
	"nena-manipu-latina/models"
)

// TODO: accommodate calling Bidders concurrently with timeout

func RunAuction(ctx context.Context, bidders []bidder.Bidder, strategyName string, req models.BidRequest) (*models.BidCollectionResult, error) {
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Starting for strategy [%s], request: %v", strategyName, req))

	stratImplementation, err := Factory(strategyName)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Run Auction [ERROR] after Factory: %v", err))
		return &models.BidCollectionResult{}, errors.New("run Auction error after Factory")
	}
	result, err := stratImplementation.AuctionWithContext(ctx, bidders, req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("Run Auction [ERROR] after Auction: %v", err))
		return &models.BidCollectionResult{}, errors.New("run Auction error after Auction")
	}
	eight_ball_logger.Info(fmt.Sprintf("Run Auction: Auction Complete. Winning Response: %v", result))
	return result, nil
}
