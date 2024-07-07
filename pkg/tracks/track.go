package tracks

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/common"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
)

// Track represents a track.
type Track struct {
	ID       uuid.UUID
	Title    string
	Duration uint // in milliseconds
	Tabs     []*tabs.Tab
	Links    []*common.Link
}

// TrackConfig modifies a Track with configuration options.
type TrackConfig func(*Track)

// WithID sets the ID of a Track.
func WithID(id uuid.UUID) TrackConfig {
	return func(a *Track) {
		a.ID = id
	}
}

// WithTabs sets the tabs of a Track.
func WithTabs(tabs []*tabs.Tab) TrackConfig {
	return func(a *Track) {
		a.Tabs = tabs
	}
}

// NewTrack instantiates a new Track.
func NewTrack(title string, duration uint, configs ...TrackConfig) *Track {
	track := &Track{ID: uuid.New(), Title: title, Duration: duration}
	for _, config := range configs {
		config(track)
	}
	return track
}
