package sources

// Operations represents operations related to sources.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing sources.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a new Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
