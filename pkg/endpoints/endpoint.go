package endpoints

// Endpoint represents a record in the 'endpoints' table.
type Endpoint struct {
	SourceID uint   `yaml:"sourceId"`
	Category string `yaml:"category"`
	Type     string `yaml:"type"`
	URL      string `yaml:"url"`
}

// NewEndpoint instantiates a new Endpoint.
func NewEndpoint(sourceId uint, category, endpointType, url string) *Endpoint {
	return &Endpoint{
		SourceID: sourceId,
		Category: category,
		Type:     endpointType,
		URL:      url,
	}
}
