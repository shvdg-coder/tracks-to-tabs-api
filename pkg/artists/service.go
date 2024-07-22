package artists

import arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"

// Operations represents all operations related to Artists.
type Operations interface {
	DatabaseOperations
	MappingOperations
	arttrk.Operations
}

// Service is responsible for managing and retrieving Artists.
type Service struct {
	DatabaseOperations
	MappingOperations
	ArtistTrackOps arttrk.Operations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations, mapping MappingOperations, artistTrack arttrk.Operations) Operations {
	return &Service{DatabaseOperations: database, MappingOperations: mapping, ArtistTrackOps: artistTrack}
}

// LinkArtistToTrack links an artist to a track.
func (s *Service) LinkArtistToTrack(artistId string, trackId string) {
	s.ArtistTrackOps.LinkArtistToTrack(artistId, trackId)
}

// GetArtistToTrackLink retrieves a link between an artist and a track.
func (s *Service) GetArtistToTrackLink(artistID string) (*arttrk.ArtistTrack, error) {
	return s.ArtistTrackOps.GetArtistToTrackLink(artistID)
}

// GetArtistToTrackLinks retrieves 'artist to track' links.
func (s *Service) GetArtistToTrackLinks(artistID ...string) ([]*arttrk.ArtistTrack, error) {
	return s.ArtistTrackOps.GetArtistToTrackLinks(artistID...)
}
