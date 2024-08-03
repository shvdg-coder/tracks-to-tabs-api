package difficulties

// Operations represents operations related to difficulties.
type Operations interface {
	DataOperations
}

// Service is responsible for managing difficulties.
type Service struct {
	DataOperations
}

// NewService instantiates a new instance of Service.
func NewService(data DataOperations) Operations {
	return &Service{DataOperations: data}
}
