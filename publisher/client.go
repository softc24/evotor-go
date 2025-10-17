package publisher

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/capcom6/go-restkit"
	"github.com/softc24/evotor-go"
	"github.com/softc24/evotor-go/helpers"
)

type Client struct {
	*restkit.Client

	headers map[string]string
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
		headers: map[string]string{
			"Authorization": "Bearer " + cfg.Token,
			"User-Agent":    "evotor-go/dev",
			"Accept":        "application/vnd.evotor.v2+json",
		},
	}, nil
}

func (p *Client) GetEvents(
	ctx context.Context,
	appID string,
	eventTypes []EventType,
	opts ...GetEventsOption,
) (iter.Seq[Event[any]], func() error) {
	// Validate event types
	if len(eventTypes) == 0 {
		return helpers.EmptyIter[Event[any]](), helpers.ErrorFunc(
			fmt.Errorf("%w: no event types specified", evotor.ErrBadRequest),
		)
	}

	// Apply options
	options := new(getEventsOptions)
	options.apply(opts...)

	// Build query parameters
	params := options.toQuery()

	// Join valid event types with commas
	typeStrings := make([]string, len(eventTypes))
	for i, eventType := range eventTypes {
		typeStrings[i] = string(eventType)
	}
	params.Add("type", strings.Join(typeStrings, ","))

	return helpers.GetPagedData[Event[any]](
		ctx,
		p.Client,
		fmt.Sprintf("/api/apps/%s/events", appID),
		params,
		p.headers,
	)
}
