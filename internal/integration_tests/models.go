package integration_tests

import "github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"

// ExpectedArtist contains the data of what a models.Artist is expected to have.
type ExpectedArtist struct {
	*models.ArtistEntry
	TrackCount     int
	ReferenceCount int
	ResourceCount  int
}

// ExpectedTrack contains the data of what a models.Track is expected to have.
type ExpectedTrack struct {
	*models.TrackEntry
	TabCount       int
	ReferenceCount int
	ResourceCount  int
}

// ExpectedTab contains the data of what a models.Tab is expected to have.
type ExpectedTab struct {
	*models.TabEntry
	ReferencesCount int
	ResourceCount   int
}
