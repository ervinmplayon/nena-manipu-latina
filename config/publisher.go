package config

import (
	"nena-manipu-latina/auction"
	"nena-manipu-latina/bidder"
)

// ? PublisherConfig is loaded per incoming ad request
type PublisherConfig struct {
	Strategy string
	Bidders  []bidder.Bidder
}

// ? Fake static config mapping (can later be replaced by DB or Redis)
var publisherConfigs = map[string]PublisherConfig{
	"publisher_firstresponder": {
		Strategy: auction.FirstResponderStrategy,
		Bidders:  bidder.Factory("publisher_firstresponder"),
	},
	"publisher_roundrobin": {
		Strategy: auction.RoundRobinStrategy,
		Bidders:  bidder.Factory("publisher_roundrobin"),
	},
}

// ? This returns config based on publisher ID
func GetPublisherConfig(publisherID string) (PublisherConfig, bool) {
	config, exists := publisherConfigs[publisherID]
	if !exists {
		config, exists = publisherConfigs["default"]
	}
	return config, exists
}
