package users

// Operations represents the operations related to users.
type Operations interface {
	DataOperations
}

// Service is responsible for managing users.
type Service struct {
	DataOperations
}

// NewService creates a new instance of Service
func NewService(database DataOperations) Operations {
	return &Service{DataOperations: database}
}
