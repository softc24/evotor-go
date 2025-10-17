package publisher

import (
	"net/url"
	"strconv"
)

// getEventsOptions contains optional parameters for GetEvents method.
type getEventsOptions struct {
	since int64
	until int64
	limit int
}

func (o *getEventsOptions) apply(opts ...GetEventsOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// toQuery converts getEventsOptions to url.Values.
func (o *getEventsOptions) toQuery() url.Values {
	params := url.Values{}
	if o.since != 0 {
		params.Set("since", strconv.FormatInt(o.since, 10))
	}
	if o.until != 0 {
		params.Set("until", strconv.FormatInt(o.until, 10))
	}
	if o.limit != 0 {
		params.Set("limit", strconv.Itoa(o.limit))
	}
	return params
}

// GetEventsOption is a function type that modifies GetEventsOptions.
type GetEventsOption func(*getEventsOptions)

// WithSince sets the since parameter for GetEvents.
func WithSince(since int64) GetEventsOption {
	return func(opts *getEventsOptions) {
		opts.since = since
	}
}

// WithUntil sets the until parameter for GetEvents.
func WithUntil(until int64) GetEventsOption {
	return func(opts *getEventsOptions) {
		opts.until = until
	}
}

// WithLimit sets the limit parameter for GetEvents.
func WithLimit(limit int) GetEventsOption {
	return func(opts *getEventsOptions) {
		opts.limit = limit
	}
}
