package endpoints

// Endpoint represents a record in the 'endpoints' table.
type Endpoint struct {
	SourceID int    `yaml:"sourceId"`
	Category string `yaml:"category"`
	URL      string `yaml:"url"`
}

// NewEndpoint instantiates a new Endpoint.
func NewEndpoint(sourceId int, category, url string) *Endpoint {
	return &Endpoint{
		SourceID: sourceId,
		Category: category,
		URL:      url,
	}
}
