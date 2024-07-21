package artisttrack

// Service is responsible for managing and retrieving 'artist to track' links.
type Service struct {
	*DatabaseService
}

// NewService instantiates a Service.
func NewService(database *DatabaseService) *Service {
	return &Service{DatabaseService: database}
}
