package models

import (
	"encoding/json"
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

// MarshalJSON marshals the models.Endpoint while preventing cyclic references.
func (e *Endpoint) MarshalJSON() ([]byte, error) {
	endpoint := *e
	endpoint.Source = &Source{
		SourceEntry: e.Source.SourceEntry,
		Endpoints:   nil,
	}
	return json.Marshal(endpoint)
}
