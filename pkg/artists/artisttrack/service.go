package artisttrack

import "github.com/google/uuid"

// Operations represents all operations related to 'artist to track' links.
type Operations interface {
	DatabaseOperations
	ExtractTrackIDs(artistTracks []*ArtistTrack) []uuid.UUID
}

// Service is responsible for managing and retrieving 'artist to track' links.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}

// ExtractTrackIDs retrieves the track IDs from each ArtistTrack.
func (s *Service) ExtractTrackIDs(artistTracks []*ArtistTrack) []uuid.UUID {
	var trackIDs []uuid.UUID
	for _, link := range artistTracks {
		trackIDs = append(trackIDs, link.TrackID)
	}
	return trackIDs
}
