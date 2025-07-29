package auction

import (
	"fmt"
)

var Factory = func(name string) (AuctionStrategy, error) {
	switch name {
	case RoundRobinStrategy:
		eight_ball_logger.Info(fmt.Sprintf("Auction Strategy Factory: Selected %s", RoundRobinStrategy))
		return &RoundRobin{}, nil
	// case FirstResponderStrategy:
	// 	eight_ball_logger.Info(fmt.Sprintf("Auction Strategy Factory: Selected %s", FirstResponderStrategy))
	// 	return &FirstResponder{}, nil
	default:
		err := fmt.Errorf("unknown Auction strategy: %s", name)
		eight_ball_logger.Error(fmt.Sprintf("Strategy Factory [ERROR]: %s", err))
		return nil, err
	}
}
