package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistOps represents all operations related to artists.
type ArtistOps interface {
	data.ArtistData
	mappers.ArtistMapper
	GetArtistsCascading(artistID ...uuid.UUID) ([]*models.ArtistEntry, error)
}

// ArtistSvc is responsible for managing and retrieving artists.
type ArtistSvc struct {
	data.ArtistData
	mappers.ArtistMapper
	ArtistTrackOps ArtistTrackOps
	TrackOps       TrackOps
	ReferenceOps   ReferenceOps
}

// NewArtistSvc instantiates a ArtistSvc.
func NewArtistSvc(data data.ArtistData, mapper mappers.ArtistMapper, artistTracks ArtistTrackOps, tracks TrackOps, references ReferenceOps) ArtistOps {
	return &ArtistSvc{
		ArtistData:     data,
		ArtistMapper:   mapper,
		ArtistTrackOps: artistTracks,
		TrackOps:       tracks,
		ReferenceOps:   references,
	}
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (a *ArtistSvc) GetArtistsCascading(artistID ...uuid.UUID) ([]*models.ArtistEntry, error) {
	return nil, nil
}
