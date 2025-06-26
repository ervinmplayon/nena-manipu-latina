package auction

// ? This allows future expansion via admin UI
type AuctionConfig struct {
	StrategyName string `json:"strategy"`
	// ? Later on, add TimeoutMS, MaxBidders, Weightingm etc.
}

// ! Example only
// TODO: source this from JSON, environment or DB later
var CurrentAuctionConfig = AuctionConfig{
	StrategyName: FirstResponderStrategy, // or RoundRobinStrategy
}
