package evotor

import "net/http"

// Config contains client options.
type Config struct {
	BaseURL string

	Client *http.Client
}
