package models

import "encoding/json"

// DifficultyEntry represents a difficulty level in the database.
type DifficultyEntry struct {
	ID   uint   `yaml:"id" db:"id"`
	Name string `yaml:"name" db:"name"`
}

// Difficulty represents a difficulty with entity references
type Difficulty struct {
	*DifficultyEntry
}

// MarshalJSON marshals the models.Difficulty.
func (d *Difficulty) MarshalJSON() ([]byte, error) {
	return json.Marshal(*d)
}
