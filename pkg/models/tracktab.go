package models

import "github.com/google/uuid"

// TrackTabEntry represents a link from a track to a tab.
type TrackTabEntry struct {
	TrackID uuid.UUID
	TabID   uuid.UUID
}

// TrackTab represents a 'track to tab' link with entity references.
type TrackTab struct {
	Track *Track
	Tab   *Tab
}
