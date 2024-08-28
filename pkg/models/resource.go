package models

import (
	"encoding/json"
	"strings"
)

// Resource represents a data item of a particular source.
type Resource struct {
	FormattedURL string
	*Endpoint
}

// FormatURL formats the URL with the corresponding values.
func (r *Resource) FormatURL(replacements map[string]string) {
	url := r.UnformattedURL
	for placeholder, replacement := range replacements {
		url = strings.Replace(url, placeholder, replacement, -1)
	}
	r.FormattedURL = url
}

// MarshalJSON marshals the models.Resource.
func (r *Resource) MarshalJSON() ([]byte, error) {
	resource := *r
	resource.Source = nil
	return json.Marshal(*r)
}
