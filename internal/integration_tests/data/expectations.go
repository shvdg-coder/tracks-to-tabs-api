package dataAPI

import (
	"github.com/google/uuid"
	"testing"
)

// createExpectedArtists constructs and returns a map of ExpectedArtist's for use in test cases.
func createExpectedArtists(t *testing.T) map[uuid.UUID]*ExpectedArtist {
	expectedArtists := make(map[uuid.UUID]*ExpectedArtist)

	artistsMap := createArtistsFromCSV(t, artistsCSV)
	for id, artist := range artistsMap {
		expectedArtist := &ExpectedArtist{
			ArtistEntry: artist,
		}
		expectedArtists[id] = expectedArtist
	}

	return expectedArtists
}

// createExpectedTracks constructs and returns a map of ExpectedTrack's for use in test cases.
func createExpectedTracks(t *testing.T) map[uuid.UUID]*ExpectedTrack {
	expectedTracks := make(map[uuid.UUID]*ExpectedTrack)

	tracksMap := createTracksFromCSV(t, tracksCSV)
	for id, track := range tracksMap {
		expectedTrack := &ExpectedTrack{
			TrackEntry: track,
		}
		expectedTracks[id] = expectedTrack
	}

	return expectedTracks
}

// createExpectedTabs constructs and returns a map of ExpectedTab's for use in test cases.
func createExpectedTabs(t *testing.T) map[uuid.UUID]*ExpectedTab {
	expectedTabs := make(map[uuid.UUID]*ExpectedTab)

	tabsMap := createTabsFromCSV(t, tabsCSV)
	for _, tab := range tabsMap {
		expectedTab := &ExpectedTab{
			TabEntry: tab,
		}
		expectedTabs[tab.ID] = expectedTab
	}

	return expectedTabs
}
