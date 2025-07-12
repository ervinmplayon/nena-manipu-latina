// ? Decorator that logs every bid
package bidder

import (
	"errors"
	"fmt"
	"nena-manipu-latina/models"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

type ChismosaBidder struct {
	Inner  Bidder
	Logger logger.Logger
}

// TODO: Properly implement bidder interface. Implement Name()

func (cb *ChismosaBidder) Bid(req models.BidRequest) (models.BidResponse, error) {
	cb.Logger.Info("Processing bid request")
	res, err := cb.Inner.Bid(req)
	if err != nil {
		cb.Logger.Error(fmt.Sprintf("ChismeBidder Error: %s", err))
		return models.BidResponse{}, errors.New("chismosa Bidder Error")
	}
	cb.Logger.Info("Returning bid response")
	return res, nil
}
