package bidder

import (
	"net/http"
	"time"
)

type HTTPBidder struct {
	name     string
	endpoint string
	timeout  time.Duration
	client   *http.Client
}

func NewHTTPBidder(name, endpoint string, timeout time.Duration) *HTTPBidder {
	return &HTTPBidder{
		name:     name,
		endpoint: endpoint,
		timeout:  timeout,
		client:   &http.Client{},
	}
}

func (hb *HTTPBidder) Name() string {
	return hb.name
}
