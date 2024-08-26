package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// TabEntry represents a tab.
type TabEntry struct {
	ID           uuid.UUID
	InstrumentID uint
	DifficultyID uint
	Description  string
}

// Tab represents a tab.
type Tab struct {
	*TabEntry
	Track      *Track
	Instrument *Instrument
	Difficulty *Difficulty
	References []*Reference
}

// MarshalJSON marshals the models.Tab while preventing circling.
func (t *Tab) MarshalJSON() ([]byte, error) {
	tab := *t
	tab.Track = &Track{
		TrackEntry: t.Track.TrackEntry,
		Artist:     t.Track.Artist,
	}
	return json.Marshal(tab)
}
