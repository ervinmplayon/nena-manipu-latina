// ? Decorator that logs every bid
package bidder

import (
	"fmt"
	"nena-manipu-latina/models"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

type ChismosaBidder struct {
	Inner  Bidder
	Logger logger.Logger
}

func (cb *ChismosaBidder) Bid(req models.AdRequest) models.AdResponse {
	cb.Logger.Info("Processing bid request")
	res, err := cb.Inner.Bid(req)
	if err != nil {
		cb.Logger.Error(fmt.Sprintf("ChismeBidder Error: %s", err))
		return models.AdResponse{}
	}
	cb.Logger.Info("Returning bid response")
	return res
}
