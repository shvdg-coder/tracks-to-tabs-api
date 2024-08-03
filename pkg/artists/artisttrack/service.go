package artisttrack

import "github.com/google/uuid"

// Operations represents all operations related to 'artist to track' links.
type Operations interface {
	DataOperations
	ExtractTrackIDs(artistTracks []*ArtistTrack) []uuid.UUID
}

// Service is responsible for managing and retrieving 'artist to track' links.
type Service struct {
	DataOperations
}

// NewService instantiates a Service.
func NewService(data DataOperations) Operations {
	return &Service{DataOperations: data}
}

// ExtractTrackIDs retrieves the track IDs from each ArtistTrack.
func (s *Service) ExtractTrackIDs(artistTracks []*ArtistTrack) []uuid.UUID {
	var trackIDs []uuid.UUID
	for _, link := range artistTracks {
		trackIDs = append(trackIDs, link.TrackID)
	}
	return trackIDs
}
