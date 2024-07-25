package users

// Operations represents the operations related to users.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing users.
type Service struct {
	DatabaseOperations
}

// NewService creates a new instance of Service
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
