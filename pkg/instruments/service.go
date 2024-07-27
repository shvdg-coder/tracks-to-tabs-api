package instruments

// Operations represents the operations related to instruments.
type Operations interface {
	DataOperations
}

// Service is responsible for managing instruments.
type Service struct {
	DataOperations
}

// NewService creates a new instance of Service.
func NewService(database DataOperations) Operations {
	return &Service{DataOperations: database}
}
