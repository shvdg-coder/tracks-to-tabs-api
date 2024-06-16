package tabs

import (
	"github.com/google/uuid"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
)

// Tab represents a tab.
type Tab struct {
	ID          uuid.UUID
	Instrument  *inst.Instrument
	Difficulty  *diff.Difficulty
	Description string
}

// NewTab instantiates a new Tab.
func NewTab(instrument *inst.Instrument, difficulty *diff.Difficulty, description string) *Tab {
	return &Tab{
		ID:          uuid.New(),
		Instrument:  instrument,
		Difficulty:  difficulty,
		Description: description,
	}
}
