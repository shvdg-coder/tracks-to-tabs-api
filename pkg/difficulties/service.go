package difficulties

// Operations represents operations related to difficulties.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing difficulties.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a new instance of Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
