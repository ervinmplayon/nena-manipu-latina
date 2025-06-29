package auction

import (
	"fmt"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

// ? This is declared at the package level, it has package scope. It can be used anywhere in this package
var eight_ball_logger logger.Logger = &logger.StandardLogger{}

func Factory(name string) (AuctionStrategy, error) {
	switch name {
	case RoundRobinStrategy:
		eight_ball_logger.Info(fmt.Sprintf("Auction Strategy Factory: Selected %s", RoundRobinStrategy))
		return &RoundRobin{}, nil
	case FirstResponderStrategy:
		eight_ball_logger.Info(fmt.Sprintf("Auction Strategy Factory: Selected %s", FirstResponderStrategy))
		return &FirstResponder{}, nil
	default:
		err := fmt.Errorf("unknown Auction strategy: %s", name)
		eight_ball_logger.Error(fmt.Sprintf("Strategy Factory [ERROR]: %s", err))
		return nil, err
	}
}
