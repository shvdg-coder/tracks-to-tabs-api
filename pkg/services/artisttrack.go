package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
)
import (
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/schemas"
)

// ArtistTrackOps combines the interface representing operations related to 'artist to track' links.
type ArtistTrackOps interface {
	schemas.ArtistTrackSchema
	data.ArtistTrackData
	CreateArtistTrackEntries(artist *models.ArtistEntry, tracks []*models.TrackEntry) []*models.ArtistTrackEntry
	ExtractIDsFromArtistTrackEntries(artistTracks []*models.ArtistTrackEntry) (artistIDs []uuid.UUID, trackIDs []uuid.UUID)
}

// ArtistTrackSvc is responsible for managing 'artist to track' links.
type ArtistTrackSvc struct {
	schemas.ArtistTrackSchema
	data.ArtistTrackData
}

// NewArtistTrackSvc instantiates an ArtistTrackSvc.
func NewArtistTrackSvc(schema schemas.ArtistTrackSchema, data data.ArtistTrackData) ArtistTrackOps {
	return &ArtistTrackSvc{
		ArtistTrackSchema: schema,
		ArtistTrackData:   data,
	}
}

// CreateArtistTrackEntries creates models.ArtistTrack's using the provided models.ArtistEntry and models.TrackEntry's.
func (at *ArtistTrackSvc) CreateArtistTrackEntries(artist *models.ArtistEntry, tracks []*models.TrackEntry) []*models.ArtistTrackEntry {
	artistTracks := make([]*models.ArtistTrackEntry, len(tracks))
	for i, track := range tracks {
		artistTracks[i] = &models.ArtistTrackEntry{ArtistID: artist.ID, TrackID: track.ID}
	}
	return artistTracks
}

// ExtractIDsFromArtistTrackEntries retrieves the artist and track IDs from the models.ArtistTrackEntry's.
func (at *ArtistTrackSvc) ExtractIDsFromArtistTrackEntries(artistTracks []*models.ArtistTrackEntry) (artistIDs []uuid.UUID, trackIDs []uuid.UUID) {
	artistIDs = make([]uuid.UUID, len(artistTracks))
	trackIDs = make([]uuid.UUID, len(artistTracks))

	for _, artistTrack := range artistTracks {
		artistIDs = append(artistIDs, artistTrack.ArtistID)
		trackIDs = append(trackIDs, artistTrack.TrackID)
	}

	return artistIDs, trackIDs
}
