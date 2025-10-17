package publisher

import (
	"net/url"
	"strconv"
	"time"
)

// getEventsOptions contains optional parameters for GetEvents method.
type getEventsOptions struct {
	since int64
	until int64
	limit uint32
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
		params.Set("limit", strconv.FormatUint(uint64(o.limit), 10))
	}
	return params
}

// GetEventsOption is a function type that modifies GetEventsOptions.
type GetEventsOption func(*getEventsOptions)

// WithSince sets the since parameter for GetEvents.
func WithSince(since int64) GetEventsOption {
	return func(opts *getEventsOptions) {
		opts.since = max(since, 0)
	}
}

// WithSinceTime sets the since parameter for GetEvents.
func WithSinceTime(since time.Time) GetEventsOption {
	return WithSince(since.UnixMilli())
}

// WithUntil sets the until parameter for GetEvents.
func WithUntil(until int64) GetEventsOption {
	return func(opts *getEventsOptions) {
		opts.until = max(until, 0)
	}
}

// WithUntilTime sets the until parameter for GetEvents.
func WithUntilTime(until time.Time) GetEventsOption {
	return WithUntil(until.UnixMilli())
}

// WithLimit sets the limit parameter for GetEvents.
func WithLimit(limit uint32) GetEventsOption {
	return func(opts *getEventsOptions) {
		opts.limit = limit
	}
}
