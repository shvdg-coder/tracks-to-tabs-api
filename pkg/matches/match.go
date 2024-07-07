package matches

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

type ArtistMatch struct {
	Artist *artists.Artist
	Links  []*Link
}

type TrackMatch struct {
	Track *tracks.Track
	Links []*Link
}

type TabMatch struct {
	Tab  *tabs.Tab
	Link *Link
}

type Match struct {
	Artist *ArtistMatch
	Track  *TrackMatch
	Tabs   []*TabMatch
}
