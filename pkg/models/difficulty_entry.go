package models

// DifficultyEntry represents a difficulty level in the database.
type DifficultyEntry struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}
