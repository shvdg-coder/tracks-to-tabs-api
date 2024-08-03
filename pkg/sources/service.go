package sources

// Operations represents operations related to sources.
type Operations interface {
	DataOperations
}

// Service is responsible for managing sources.
type Service struct {
	DataOperations
}

// NewService instantiates a new Service.
func NewService(data DataOperations) Operations {
	return &Service{DataOperations: data}
}
