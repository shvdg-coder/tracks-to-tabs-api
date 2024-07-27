package tabs

// Operations represent all operations related to tabs.
type Operations interface {
	DataOperations
	MappingOperations
}

// Service is responsible for managing and retrieving tabs.
type Service struct {
	DataOperations
	MappingOperations
}

// NewService instantiates a Service.
func NewService(database DataOperations, mapping MappingOperations) Operations {
	return &Service{
		DataOperations:    database,
		MappingOperations: mapping,
	}
}
