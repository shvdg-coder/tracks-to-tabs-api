package difficulties

// Difficulty represents a difficulty level.
type Difficulty struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// Option modifies a Difficulty with configuration options.
type Option func(*Difficulty)

// WithID sets the ID of a Difficulty.
func WithID(id uint) Option {
	return func(d *Difficulty) {
		d.ID = id
	}
}

// NewDifficulty instantiates a new Difficulty.
func NewDifficulty(name string, options ...Option) *Difficulty {
	difficulty := &Difficulty{Name: name}
	for _, option := range options {
		option(difficulty)
	}
	return difficulty
}
