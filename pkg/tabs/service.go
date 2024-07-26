package tabs

// Operations represent all operations related to tabs.
type Operations interface {
	DatabaseOperations
	MappingOperations
}

// Service is responsible for managing and retrieving tabs.
type Service struct {
	DatabaseOperations
	MappingOperations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations, mapping MappingOperations) Operations {
	return &Service{
		DatabaseOperations: database,
		MappingOperations:  mapping,
	}
}
