package models

type BidCollectionResult struct {
	Responses []*BidResponse
	Errors    []error
	Metrics   BidCollectionMetrics
}
