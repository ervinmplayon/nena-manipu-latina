// ? Assumes bids are collected in order of arrival (concurrent fetch)
// ? Picks the first valid bid
package auction

import "time"

type FirstResponder struct {
	Timeout          time.Duration
	PerBidderTimeout time.Duration
}

func NewFirstResponder(timeout, perBidderTimeout time.Duration) *FirstResponder {
	return &FirstResponder{
		Timeout:          timeout,
		PerBidderTimeout: perBidderTimeout,
	}
}
