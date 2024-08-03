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
func NewService(data DataOperations, mapping MappingOperations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
	}
}
