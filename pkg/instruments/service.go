package instruments

// Operations represents the operations related to instruments.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing instruments.
type Service struct {
	DatabaseOperations
}

// NewService creates a new instance of Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
