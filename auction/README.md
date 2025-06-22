# Temporary Tattoos
At a high level, production-grade ad servers typically switch auction strategies in one of 3 ways, depending on the flexibility, scale and runtime control.

## Environment Configs at Startup
* Most common for low-complexity, high-stability environments
* Strategy is determined by an environment variable like `AUCTION_STRATEGY=highest_bid`
* On startup, the server reads this value and instantiates the corresponding strategy
### Pros
* Easy to manage and deploy
* No runtime performance cost
### Cons
* Requires restart to change behavior
### In a nutshell
* No Runtime switch
* Global granularity
* For simple deployments

## Runtime Configs via Feature Flags/Admin UI
* Uses a remote config service or a DB-backed config table (e.g., `auction_strategy` column per account or campaign)
* Strategy is read at runtime per request or cached periodically.
### Pros
* Dynamically switch strategies without restarts
* Can do A/B testing or per-campaign customization
### Cons
* Slightly more complex, config read layer must be active and reliable
* Needs strategy fallbacks if unknown
### In a nutshell
* Has Runtime switch
* Per-request/user granularity
* Good for use cases with A/B tests, feature flags

## Dynamic Bidding per Publisher/Campaign
* Auction strategy is part of the campaign metadata
* For example, a DSP might support: `strategy: "highest_bid"`, `strategy: "weighted_random"`, `strategy: "second_price"` (second price is for programmatic guaranteed deals)
### Pros
* Per-request flexibility
* Campaign-specific bidding logic
### Cons
* Increases complexity in the auction runner
* might require caching for performance
### In a nutshell
* Has Runtime switch
* Campaign-specific granularity
* Use case for advanced programmatic setups

## Example of switching strategies via config or env vars, define a factory.
```go
func GetStrategy(name string) auction.AuctionStrategy {
	switch name {
	case "highest":
		return &auction.HighestBidWins{}
	case "roundrobin":
		return &auction.RoundRobin{}
	case "first":
		return &auction.FirstResponder{}
	default:
		return &auction.HighestBidWins{}
	}
}
```