package cloud

import "net/http"

// ClientConfig contains client options.
type ClientConfig struct {
	BaseURL string

	Client *http.Client
}
