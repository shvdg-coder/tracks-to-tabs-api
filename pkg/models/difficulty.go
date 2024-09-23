package models

import "encoding/json"

// DifficultyEntry represents a difficulty level in the database.
type DifficultyEntry struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// Fields returns a slice of interfaces containing values of the DifficultyEntry.
func (d *DifficultyEntry) Fields() []interface{} {
	return []interface{}{d.ID, d.Name}
}

// Difficulty represents a difficulty with entity references
type Difficulty struct {
	*DifficultyEntry
}

// MarshalJSON marshals the models.Difficulty.
func (d *Difficulty) MarshalJSON() ([]byte, error) {
	return json.Marshal(*d)
}
