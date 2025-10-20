package cloud

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"

	"github.com/capcom6/go-restkit"
	"github.com/softc24/evotor-go/helpers"
)

type Client struct {
	*restkit.Client
}

func NewClient(cfg ClientConfig) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultURL
	}

	rest, err := restkit.NewClient(restkit.Config{
		Client:  cfg.Client,
		BaseURL: cfg.BaseURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &Client{
		Client: rest,
	}, nil
}

// GetDevices returns a sequence of devices and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of devices, and the sequence will
// yield each device in the list. If there are more devices available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the devices, the error function will
// return the error.
func (c *Client) GetDevices(ctx context.Context, token string) (iter.Seq[Device], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	return helpers.GetPagedData[Device](
		ctx,
		c.Client,
		"/devices",
		params,
		headers,
	)
}
