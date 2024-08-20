package artists

import (
	"github.com/google/uuid"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
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
	ReferenceOps   ref.Operations
}

// NewService instantiates a Service.
func NewService(data DataOperations, mapping MappingOperations, artistTracks arttrk.Operations, tracks trk.Operations, references ref.Operations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
		ArtistTrackOps:    artistTracks,
		TrackOps:          tracks,
		ReferenceOps:      references,
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

	references, err := s.ReferenceOps.GetReferencesCascading(artistID...)
	if err != nil {
		return nil, err
	}

	artistsMap := s.ArtistsToMap(artists)
	tracksMap := s.TrackOps.TracksToMap(tracks)
	artists = s.MapTracksToArtists(artistTracks, artistsMap, tracksMap)
	artists = s.MapReferencesToArtists(artistsMap, references)

	return artists, nil
}
