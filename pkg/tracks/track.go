package tracks

import "github.com/google/uuid"

// Track represents an artist
type Track struct {
	ID       uuid.UUID
	Title    string
	Duration uint // in milliseconds
}

// NewTrack instantiates a new Track.
func NewTrack(title string, duration uint) *Track {
	return &Track{
		ID:       uuid.New(),
		Title:    title,
		Duration: duration,
	}
}
