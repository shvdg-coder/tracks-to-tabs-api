package models

import "github.com/google/uuid"

// TrackTabEntry represents a link from a track to a tab.
type TrackTabEntry struct {
	TrackID uuid.UUID `db:"track_id"`
	TabID   uuid.UUID `db:"tab_id"`
}

// Fields returns a slice of interfaces containing values of the TrackTabEntry.
func (tt *TrackTabEntry) Fields() []interface{} {
	return []interface{}{tt.TrackID, tt.TabID}
}

// TrackTab represents a 'track to tab' link with entity references.
type TrackTab struct {
	Track *Track
	Tab   *Tab
}
