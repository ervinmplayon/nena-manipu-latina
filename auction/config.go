package auction

// ? This allows future expansion via admin UI
type AuctionConfig struct {
	StrategyName string `json:"strategy"`
	// ? Later on, add TimeoutMS, MaxBidders, Weightingm etc.
}

// * This config is no longer used. Instead config/runtimeconfig is now in charge of configs
var CurrentAuctionConfig = AuctionConfig{
	StrategyName: FirstResponderStrategy, // or RoundRobinStrategy
}
