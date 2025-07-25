package models

type AuctionResult struct {
	BidderName string
	Response   *BidResponse
	Err        error
}
