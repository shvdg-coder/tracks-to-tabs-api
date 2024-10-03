package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// TabEntry represents a tab.
type TabEntry struct {
	ID           uuid.UUID `db:"id"`
	InstrumentID uint      `db:"instrument_id"`
	DifficultyID uint      `db:"difficulty_id"`
	Description  string    `db:"description"`
}

// Fields returns a slice of interfaces containing values of the TabEntry.
func (t *TabEntry) Fields() []interface{} {
	return []interface{}{t.ID, t.InstrumentID, t.DifficultyID, t.Description}
}

// Tab represents a tab.
type Tab struct {
	*TabEntry
	Track      *Track
	Instrument *Instrument
	Difficulty *Difficulty
	References []*Reference
	Resources  []*Resource
}

// MarshalJSON marshals the models.Tab while preventing circling.
func (t *Tab) MarshalJSON() ([]byte, error) {
	tab := *t
	tab.Track = &Track{
		TrackEntry: t.Track.TrackEntry,
		Artist:     t.Track.Artist,
		References: t.Track.References,
		Resources:  t.Track.Resources,
		Tabs:       nil,
	}
	return json.Marshal(tab)
}
