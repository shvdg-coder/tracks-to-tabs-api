package artists

import (
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// MappingOperations represents operations related to data mapping.
type MappingOperations interface {
	GetArtistsCascading(artistID ...string) ([]*Artist, error)
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

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (m *MappingService) GetArtistsCascading(artistID ...string) ([]*Artist, error) {
	return nil, nil
}
