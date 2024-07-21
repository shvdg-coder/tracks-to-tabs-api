package artists

import arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"

// Service is responsible for managing and retrieving Artists.
type Service struct {
	*DatabaseService
	*MappingService
	*arttrk.Service
}

// NewService instantiates a Service.
func NewService(database *DatabaseService, mapping *MappingService, artistTrack *arttrk.Service) *Service {
	return &Service{DatabaseService: database, MappingService: mapping, Service: artistTrack}
}
