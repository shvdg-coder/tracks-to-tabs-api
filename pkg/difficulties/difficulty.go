package difficulties

// Difficulty represents a difficulty level.
type Difficulty struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// DifficultyConfig modifies a Difficulty with configuration options.
type DifficultyConfig func(*Difficulty)

// WithID sets the ID of a Difficulty.
func WithID(id uint) DifficultyConfig {
	return func(a *Difficulty) {
		a.ID = id
	}
}

// NewDifficulty instantiates a new Difficulty.
func NewDifficulty(name string, configs ...DifficultyConfig) *Difficulty {
	instrument := &Difficulty{Name: name}
	for _, config := range configs {
		config(instrument)
	}
	return instrument
}
