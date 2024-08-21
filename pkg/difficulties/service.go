package difficulties

// Operations represents operations related to difficulties.
type Operations interface {
	DataOperations
	DifficultiesToMap(difficulties []*Difficulty) map[uint]*Difficulty
}

// Service is responsible for managing difficulties.
type Service struct {
	DataOperations
}

// NewService instantiates a new instance of Service.
func NewService(data DataOperations) Operations {
	return &Service{DataOperations: data}
}

// DifficultiesToMap transforms a slice of Difficulty's into a map, where the key is the ID and the value the Difficulty.
func (s *Service) DifficultiesToMap(difficulties []*Difficulty) map[uint]*Difficulty {
	difficultyMap := make(map[uint]*Difficulty)
	for _, difficulty := range difficulties {
		difficultyMap[difficulty.ID] = difficulty
	}
	return difficultyMap
}
