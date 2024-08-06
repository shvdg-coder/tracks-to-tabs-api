package artists

import (
	"github.com/google/uuid"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// Operations represents all operations related to artists.
type Operations interface {
	DataOperations
	MappingOperations
	GetArtistsCascading(artistID ...uuid.UUID) ([]*Artist, error)
}

// Service is responsible for managing and retrieving artists.
type Service struct {
	DataOperations
	MappingOperations
	ArtistTrackOps arttrk.Operations
	TrackOps       trk.Operations
}

// NewService instantiates a Service.
func NewService(data DataOperations, mapping MappingOperations, artistTracks arttrk.Operations, tracks trk.Operations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
		ArtistTrackOps:    artistTracks,
		TrackOps:          tracks,
	}
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (s *Service) GetArtistsCascading(artistID ...uuid.UUID) ([]*Artist, error) {
	artists, err := s.GetArtists(artistID...)
	if err != nil {
		return nil, err
	}
	artistTracks, err := s.ArtistTrackOps.GetArtistToTrackLinks(artistID...)
	if err != nil {
		return nil, err
	}
	trackIDs := s.ArtistTrackOps.ExtractTrackIDs(artistTracks)
	tracks, err := s.TrackOps.GetTracksCascading(trackIDs...)
	if err != nil {
		return nil, err
	}
	artistsMap := s.ToMap(artists)
	tracksMap := s.TrackOps.ToMap(tracks)
	artists = s.MapTracksToArtists(artistTracks, artistsMap, tracksMap)
	return artists, nil
}
