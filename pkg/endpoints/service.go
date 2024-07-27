package endpoints

// Operations represents operations related to endpoints.
type Operations interface {
	DataOperations
}

// Service is responsible for managing endpoints.
type Service struct {
	DataOperations
}

// NewService instantiates a new Service.
func NewService(database DataOperations) Operations {
	return &Service{DataOperations: database}
}
