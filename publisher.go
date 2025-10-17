package evotor

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/capcom6/go-restkit"
)

type Publisher struct {
	*restkit.Client

	headers map[string]string
}

func NewPublisher(cfg PublisherConfig) (*Publisher, error) {
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

	return &Publisher{
		Client: rest,
		headers: map[string]string{
			"Authorization": "Bearer " + cfg.Token,
			"User-Agent":    "go-evotor/dev",
			"Accept":        "application/vnd.evotor.v2+json",
		},
	}, nil
}

func (p *Publisher) GetEvents(
	ctx context.Context,
	appID string,
	eventTypes []EventType,
	opts ...GetEventsOption,
) (iter.Seq[Event[any]], func() error) {
	// Validate event types
	if len(eventTypes) == 0 {
		return func(_ func(Event[any]) bool) {}, func() error {
			return fmt.Errorf("%w: no event types specified", ErrBadRequest)
		}
	}
	for _, eventType := range eventTypes {
		if !isValidEventType(eventType) {
			return func(_ func(Event[any]) bool) {}, func() error {
				return fmt.Errorf("%w: invalid event type: %s", ErrBadRequest, eventType)
			}
		}
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

	var err error
	return func(yield func(Event[any]) bool) {
			for {
				res := new(pagedResponse[Event[any]])

				if err = p.Do(ctx, "GET", fmt.Sprintf("/api/apps/%s/events?%s", appID, params.Encode()), p.headers, nil, res); err != nil {
					return
				}

				for _, event := range res.Items {
					if !yield(event) {
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
			if err == nil {
				return nil
			}

			return fmt.Errorf("failed to get events: %w", err)
		}
}

// Helper function to validate event type.
func isValidEventType(eventType EventType) bool {
	switch eventType {
	case EventTypeDocument, EventTypeProduct, EventTypeProductGroup, EventTypeSettings, EventTypeMarketplacePurchase:
		return true
	default:
		return false
	}
}
