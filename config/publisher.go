package config

import (
	"nena-manipu-latina/auction"
	"nena-manipu-latina/bidder"
	"time"
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
		Bidders: []bidder.Bidder{
			&bidder.MockBidder{Name: "FastDSP", Delay: 20 * time.Millisecond, CPM: 1.05},
			&bidder.MockBidder{Name: "ReliableDSP", Delay: 40 * time.Millisecond, CPM: 1.10},
		},
	},
	"publisher_roundrobin": {
		Strategy: auction.RoundRobinStrategy,
		Bidders: []bidder.Bidder{
			&bidder.MockBidder{Name: "SlowBidder", Delay: 100 * time.Millisecond, CPM: 0.99},
			&bidder.MockBidder{Name: "AggressiveDSP", Delay: 50 * time.Millisecond, CPM: 1.30},
		},
	},
}

// ? This returns config based on publisher ID
func GetPublisherConfig(publisherID string) (PublisherConfig, bool) {
	config, exists := publisherConfigs[publisherID]
	return config, exists
}
