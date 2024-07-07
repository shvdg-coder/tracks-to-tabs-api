package common

import (
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"strings"
)

// Link refers to some media.
type Link struct {
	Source       *src.Source
	Endpoint     *end.Endpoint
	FormattedURL string
}

// NewLink creates a new instance of Link, by formatting the endpoint URL with the corresponding values.
func NewLink(source *src.Source, endpoint *end.Endpoint, replacements map[string]string) *Link {
	url := endpoint.URL
	for old, replacement := range replacements {
		url = strings.Replace(url, old, replacement, 1)
	}
	return &Link{Source: source, Endpoint: endpoint, FormattedURL: url}
}
