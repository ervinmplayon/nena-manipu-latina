// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import "nena-manipu-latina/models"

type FirstResponder struct{}

func (fr *FirstResponder) Auction(bids []models.AdResponse) models.AdResponse {
	if len(bids) == 0 {
		return models.AdResponse{}
	}
	return bids[0]
}
