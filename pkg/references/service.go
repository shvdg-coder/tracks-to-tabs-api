package references

// Operations represents operations related to references.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing references.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a new Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}
