package models

import (
	"encoding/json"
	"strings"
)

// EndpointEntry represents a record in the 'endpoints' table.
type EndpointEntry struct {
	SourceID       uint   `yaml:"sourceId"`
	Category       string `yaml:"category"`
	Type           string `yaml:"type"`
	UnformattedURL string `yaml:"url"`
}

// Endpoint represents an endpoint with entity references
type Endpoint struct {
	*EndpointEntry
	Source *Source
}

// CreateFormattedURL formats the endpoint UnformattedURL with the corresponding values.
func (e *EndpointEntry) CreateFormattedURL(replacements map[string]string) string {
	url := e.UnformattedURL
	for old, replacement := range replacements {
		url = strings.Replace(url, old, replacement, 1)
	}
	return url
}

// MarshalJSON marshals the models.Endpoint while preventing cyclic references.
func (e *Endpoint) MarshalJSON() ([]byte, error) {
	endpoint := *e
	endpoint.Source = &Source{
		SourceEntry: e.Source.SourceEntry,
	}
	return json.Marshal(endpoint)
}
