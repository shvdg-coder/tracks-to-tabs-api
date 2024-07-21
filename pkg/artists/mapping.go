package artists

import (
	arttrck "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	trcks "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// MappingService is responsible for mapping entities to artists.
type MappingService struct {
	ArtistAPI      *DatabaseService
	ArtistTrackAPI *arttrck.DatabaseService
	TrackAPI       *trcks.API
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService(artistAPI *DatabaseService, artistTrackAPI *arttrck.DatabaseService, tracksAPI *trcks.API) *MappingService {
	return &MappingService{ArtistAPI: artistAPI, ArtistTrackAPI: artistTrackAPI, TrackAPI: tracksAPI}
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (m *MappingService) GetArtistsCascading(artistID ...string) ([]*Artist, error) {
	return nil, nil
}
