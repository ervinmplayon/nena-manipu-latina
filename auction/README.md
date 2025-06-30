# Temporary Tattoos
# TODO
- [ ] Emphasize how we actually built an SSP good for A/B testing and feature flags.

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

## Variables declared at the Package Level 
A variable declared at the package level (oustide of any function) can be reused and accessed by any function or type within the same package. This is because package-level variables have package scope, meaning that their visibility extends across all source files belonging to that package. 
* Declaration: They are declared using the `var` keyword at the top level of a file within a package, outside of any function. 
* Scope: They are accessible throughout the entire package.
* Visibility: If the variable name begins with a lowercase letter, it is only accessible within that package (unexported). If it begins with an uppercase letter, it is exported and can be accessed by other packages that import it. 
* Lifetime: They exist for the entire duration of the program's execution. 

## Message to ID