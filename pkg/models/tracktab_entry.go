package models

import "github.com/google/uuid"

// TrackTabEntry represents a link from a track to a tab.
type TrackTabEntry struct {
	TrackID uuid.UUID
	TabID   uuid.UUID
}
