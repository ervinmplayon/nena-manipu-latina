// ? Decorator that logs every bid
package bidder

import (
	"nena-manipu-latina/models"

	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

type ChismosaBidder struct {
	Inner  Bidder
	Logger logger.Logger
}

func (cb *ChismosaBidder) Bid(req models.AdRequest) models.AdResponse {
	cb.Logger.Info("Processing bid request")
	res := cb.Inner.Bid(req)
	cb.Logger.Info("Returning bid response")
	return res
}
