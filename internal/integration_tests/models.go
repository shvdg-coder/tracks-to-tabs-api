package integration_tests

import "github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"

// ExpectedArtist contains the data of what a models.Artist is expected to have.
type ExpectedArtist struct {
	*models.ArtistEntry
	Resources []string
}

// ExpectedTrack contains the data of what a models.Track is expected to have.
type ExpectedTrack struct {
	*models.TrackEntry
}

// ExpectedTab contains the data of what a models.Tab is expected to have.
type ExpectedTab struct {
	*models.TabEntry
}
