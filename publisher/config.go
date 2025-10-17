package publisher

import "net/http"

type ClientConfig struct {
	BaseURL string
	Token   string

	Client *http.Client
}
