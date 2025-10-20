package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"strings"

	"github.com/capcom6/go-restkit"
	"github.com/softc24/evotor-go/helpers"
)

type Client struct {
	*restkit.Client

	headers http.Header
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

		headers: http.Header{
			"Authorization": []string{"Bearer " + cfg.Token},
			"User-Agent":    []string{"evotor-go/dev"},
			"Accept":        []string{"application/vnd.evotor.v2+json"},
		},
	}, nil
}

func (p *Client) GetEvents(
	ctx context.Context,
	appID string,
	eventTypes []EventType,
	opts ...GetEventsOption,
) (iter.Seq[Event[json.RawMessage]], func() error) {
	// Apply options
	options := new(getEventsOptions)
	options.apply(opts...)

	// Build query parameters
	params := options.toQuery()

	// Join event types with commas
	if len(eventTypes) > 0 {
		vals := make([]string, 0, len(eventTypes))
		for _, t := range eventTypes {
			vals = append(vals, string(t))
		}
		params.Add("type", strings.Join(vals, ","))
	}

	return helpers.GetPagedData[Event[json.RawMessage]](
		ctx,
		p.Client,
		fmt.Sprintf("/api/apps/%s/events", url.PathEscape(appID)),
		params,
		p.headers,
	)
}
