package endpoints

// Endpoint represents a record in the 'endpoints' table.
type Endpoint struct {
	SourceID int
	Category string
	Endpoint string
}

// NewEndpoint instantiates a new Endpoint.
func NewEndpoint(sourceId int, category, endpoint string) *Endpoint {
	return &Endpoint{
		SourceID: sourceId,
		Category: category,
		Endpoint: endpoint,
	}
}
