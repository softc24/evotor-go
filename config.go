package evotor

import "net/http"

// Config contains client options.
type Config struct {
	BaseURL string

	Client *http.Client
}

type PublisherConfig struct {
	BaseURL string
	Token   string

	Client *http.Client
}
