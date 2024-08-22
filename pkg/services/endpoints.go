package services

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"

// ArtistTrackOps represents operations related to endpoints.
type Operations interface {
	database.TabsOps
}

// ArtistTrackSvc is responsible for managing endpoints.
type Service struct {
	database.TabsOps
}

// NewTrackSvc instantiates a new ArtistTrackSvc.
func NewService(data database.TabsOps) Operations {
	return &Service{TabsOps: data}
}
