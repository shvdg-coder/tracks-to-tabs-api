package tabs

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/commons"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
)

// Tab represents a tab.
type Tab struct {
	ID          uuid.UUID
	Instrument  *inst.Instrument
	Difficulty  *diff.Difficulty
	Description string
	Links       []*commons.Link
}

// Option modifies a Tab with configuration options.
type Option func(*Tab)

// WithID sets the ID of a Tab.
func WithID(id uuid.UUID) Option {
	return func(a *Tab) {
		a.ID = id
	}
}

// NewTab instantiates a new Tab.
func NewTab(instrument *inst.Instrument, difficulty *diff.Difficulty, description string, configs ...Option) *Tab {
	tab := &Tab{ID: uuid.New(), Instrument: instrument, Difficulty: difficulty, Description: description}
	for _, config := range configs {
		config(tab)
	}
	return tab
}
