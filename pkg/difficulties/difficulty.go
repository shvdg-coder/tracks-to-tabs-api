package difficulties

// Difficulty represents an artist
type Difficulty struct {
	ID   uint
	Name string
}

// NewDifficulty instantiates a new Difficulty.
func NewDifficulty(id uint, name string) *Difficulty {
	return &Difficulty{
		ID:   id,
		Name: name,
	}
}
