package bidder

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"nena-manipu-latina/models"
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

func (hb *HTTPBidder) Bid(req models.BidRequest) (models.BidResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), hb.timeout)
	defer cancel()

	body, err := json.Marshal(req)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("http Bidder marshal error: %s", err.Error()))
		err_msg := fmt.Sprintf("http Bidder marshal error: %s", err.Error())
		return models.BidResponse{}, errors.New(err_msg)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", hb.endpoint, bytes.NewBuffer(body))
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("http Bidder request creation error: %s", err.Error()))
		err_msg := fmt.Sprintf("http Bidder request creation error: %s", err.Error())
		return models.BidResponse{}, errors.New(err_msg)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := hb.client.Do(httpReq)
	if err != nil {
		eight_ball_logger.Error(fmt.Sprintf("http Bidder http request error: %s", err.Error()))
		err_msg := fmt.Sprintf("http Bidder http request error: %s", err.Error())
		return models.BidResponse{}, errors.New(err_msg)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		eight_ball_logger.Error(fmt.Sprintf("http Bidder http bad response status [%s]", resp.Status))
		err_msg := fmt.Sprintf("http Bidder http bad response status [%s]", resp.Status)
		return models.BidResponse{}, errors.New(err_msg)
	}

	var bidResp models.BidResponse
	if err := json.NewDecoder(resp.Body).Decode(&bidResp); err != nil {
		eight_ball_logger.Error(fmt.Sprintf("http Bidder http decode error: %s", err.Error()))
		err_msg := fmt.Sprintf("http Bidder http decode error: %s", err.Error())
		return models.BidResponse{}, errors.New(err_msg)
	}

	return bidResp, nil
}

// TODO: Write Unit Tests for this.
// * Why? Once I learn the pattern for unit testing a request, i am even more unstoppable.
