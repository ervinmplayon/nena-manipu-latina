package models

import "time"

type BidCollectionMetrics struct {
	TotalBidders      int
	SuccessfulBids    int
	TimeoutBidders    int
	FailedBidders     int
	ResponseLatencies map[string]time.Duration
}

// ? Each `bidder.Bidder` must expose a method `Name() string` for identifying itself in logs/metrics
