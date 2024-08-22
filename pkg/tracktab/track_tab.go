package tracktab

import "github.com/google/uuid"

// TrackTab represents a link from a track to a tab.
type TrackTab struct {
	TrackID uuid.UUID
	TabID   uuid.UUID
}
