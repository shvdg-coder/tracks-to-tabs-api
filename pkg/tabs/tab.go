package tabs

import "github.com/google/uuid"

// Tabs represents an artist
type Tabs struct {
	ID           uuid.UUID
	InstrumentID uint
	DifficultyID uint
	TuningID     uint
	Description  string
}

// NewTab instantiates a new Tabs.
func NewTab(instrumentId uint, difficultyId uint, tuningId uint, description string) *Tabs {
	return &Tabs{
		ID:           uuid.New(),
		InstrumentID: instrumentId,
		DifficultyID: difficultyId,
		TuningID:     tuningId,
		Description:  description,
	}
}
