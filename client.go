package evotor

import (
	"context"
	"fmt"
	"iter"
	"net/url"

	"github.com/capcom6/go-restkit"
)

type Client struct {
	*restkit.Client
}

func NewClient(cfg Config) *Client {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultURL
	}

	return &Client{
		Client: restkit.NewClient(restkit.Config{
			Client:  cfg.Client,
			BaseURL: cfg.BaseURL,
		}),
	}
}

// GetDevices returns a sequence of devices and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of devices, and the sequence will
// yield each device in the list. If there are more devices available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the devices, the error function will
// return the error.
func (c *Client) GetDevices(ctx context.Context, token string) (iter.Seq[Device], func() error) {
	var err error
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"Accept":        "application/json",
	}
	res := new(pagedResponse[Device])
	params := url.Values{}

	return func(yield func(Device) bool) {
			for {
				if err = c.Do(ctx, "GET", "/devices?"+params.Encode(), headers, nil, res); err != nil {
					return
				}

				for _, device := range res.Items {
					if !yield(device) {
						return
					}
				}

				if res.HasNext() {
					params.Set("cursor", res.Paging.NextCursor)
				} else {
					return
				}
			}
		}, func() error {
			return fmt.Errorf("failed to load devices: %w", err)
		}
}
