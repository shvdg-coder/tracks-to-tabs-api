package artists

import (
	"github.com/google/uuid"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// MappingOperations represents operations related to artist data mapping.
type MappingOperations interface {
	GetArtistsAsMap(artistID ...uuid.UUID) (map[uuid.UUID]*Artist, error)
	GetArtistsCascading(artistID ...uuid.UUID) ([]*Artist, error)
}

// MappingService is responsible for mapping entities to artists.
type MappingService struct {
	DatabaseOperations
	ArtistTrackOps arttrk.Operations
	TracksOps      trk.Operations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService(artists DatabaseOperations, artistTrack arttrk.Operations, tracks trk.Operations) MappingOperations {
	return &MappingService{
		DatabaseOperations: artists,
		ArtistTrackOps:     artistTrack,
		TracksOps:          tracks}
}

// GetArtistsAsMap retrieves the artists for the provided IDs, and creates a map where the key is the ID and the value the Artist.
func (m *MappingService) GetArtistsAsMap(artistID ...uuid.UUID) (map[uuid.UUID]*Artist, error) {
	artists, err := m.GetArtists(artistID...)
	if err != nil {
		return nil, err
	}

	artistMap := make(map[uuid.UUID]*Artist, len(artists))
	for _, artist := range artists {
		artistMap[artist.ID] = artist
	}

	return artistMap, nil
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (m *MappingService) GetArtistsCascading(artistID ...uuid.UUID) ([]*Artist, error) {
	// Get artists
	artistsMap, err := m.GetArtistsAsMap(artistID...)
	if err != nil {
		return nil, err
	}

	// Get 'artist to track' links
	artistTrackLinks, err := m.ArtistTrackOps.GetArtistToTrackLinks(artistID...)
	if err != nil {
		return nil, err
	}

	// Pluck the trackID's
	var trackIDs []uuid.UUID
	for _, artistTrackLink := range artistTrackLinks {
		trackIDs = append(trackIDs, artistTrackLink.TrackID)
	}

	// Get the tracks
	tracksMap, err := m.TracksOps.GetTracksCascadingAsMap(trackIDs...)
	if err != nil {
		return nil, err
	}

	// Map tracks to artist
	for _, artistTrackLink := range artistTrackLinks {
		artist := artistsMap[artistTrackLink.ArtistID]
		track := tracksMap[artistTrackLink.TrackID]
		// Update artist in artistMap with new track
		artist.Tracks = append(artist.Tracks, track)
	}

	// Pluck the artists from the map
	var artists []*Artist
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}

	// Done
	return artists, nil
}
