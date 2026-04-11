package coincap

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return  &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{},
		},
	}, nil
}
