package artists

import (
	"github.com/google/uuid"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// Operations represents all operations related to artists.
type Operations interface {
	DatabaseOperations
	MappingOperations
	arttrk.Operations
	GetArtistsCascading(artistID ...uuid.UUID) ([]*Artist, error)
}

// Service is responsible for managing and retrieving artists.
type Service struct {
	DatabaseOperations
	MappingOperations
	ArtistTrackOps arttrk.Operations
	TrackOps       trk.Operations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations, mapping MappingOperations, artistTracks arttrk.Operations, tracks trk.Operations) Operations {
	return &Service{
		DatabaseOperations: database,
		MappingOperations:  mapping,
		ArtistTrackOps:     artistTracks,
		TrackOps:           tracks,
	}
}

// LinkArtistToTrack links an artist to a track.
func (s *Service) LinkArtistToTrack(artistId, trackId uuid.UUID) {
	s.ArtistTrackOps.LinkArtistToTrack(artistId, trackId)
}

// GetArtistToTrackLink retrieves a link between an artist and a track.
func (s *Service) GetArtistToTrackLink(artistID uuid.UUID) (*arttrk.ArtistTrack, error) {
	return s.ArtistTrackOps.GetArtistToTrackLink(artistID)
}

// GetArtistToTrackLinks retrieves 'artist to track' links.
func (s *Service) GetArtistToTrackLinks(artistID ...uuid.UUID) ([]*arttrk.ArtistTrack, error) {
	return s.ArtistTrackOps.GetArtistToTrackLinks(artistID...)
}

// ExtractTrackIDs retrieves the track IDs from each artisttrack.ArtistTrack.
func (s *Service) ExtractTrackIDs(artistTracks []*arttrk.ArtistTrack) []uuid.UUID {
	return s.ArtistTrackOps.ExtractTrackIDs(artistTracks)
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (s *Service) GetArtistsCascading(artistID ...uuid.UUID) ([]*Artist, error) {
	artists, err := s.GetArtists(artistID...)
	if err != nil {
		return nil, err
	}
	artistTracks, err := s.GetArtistToTrackLinks(artistID...)
	if err != nil {
		return nil, err
	}
	trackIDs := s.ExtractTrackIDs(artistTracks)
	tracks, err := s.TrackOps.GetTracksCascading(trackIDs...)
	if err != nil {
		return nil, err
	}
	artistsMap := s.ToMap(artists)
	tracksMap := s.TrackOps.ToMap(tracks)
	artists = s.MapTracksToArtists(artistTracks, artistsMap, tracksMap)
	return artists, nil
}
