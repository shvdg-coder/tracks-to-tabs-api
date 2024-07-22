package artisttrack

// Operations represents all operations related to 'artist to track' links.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing and retrieving 'artist to track' links.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
