package tracktab

// Operations represents all operations related to 'track to tab' links.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing and retrieving 'track to tab' links.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
