package endpoints

// Operations represents operations related to endpoints.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing endpoints.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a new Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
