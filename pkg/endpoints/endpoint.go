package endpoints

import "strings"

// Endpoint represents a record in the 'endpoints' table.
type Endpoint struct {
	SourceID       uint   `yaml:"sourceId"`
	Category       string `yaml:"category"`
	Type           string `yaml:"type"`
	UnformattedURL string `yaml:"url"`
}

// CreateLink formats the endpoint UnformattedURL with the corresponding values.
func (e *Endpoint) CreateLink(replacements map[string]string) string {
	url := e.UnformattedURL
	for old, replacement := range replacements {
		url = strings.Replace(url, old, replacement, 1)
	}
	return url
}
