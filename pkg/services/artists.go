package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents all operations related to artists.
type Operations interface {
	database.ArtistOps
	mappers.ArtistMapper
	GetArtistsCascading(artistID ...uuid.UUID) ([]*models.ArtistEntry, error)
}

// ArtistTrackSvc is responsible for managing and retrieving artists.
type Service struct {
	database.ArtistOps
	mappers.ArtistMapper
	ArtistTrackOps Operations
	TrackOps       Operations
	ReferenceOps   Operations
}

// NewTrackSvc instantiates a ArtistTrackSvc.
func NewService(data database.ArtistOps, mapping mappers.ArtistMapper, artistTracks Operations, tracks Operations, references Operations) Operations {
	return &Service{
		ArtistOps:      data,
		ArtistOps:      mapping,
		ArtistTrackOps: artistTracks,
		TrackOps:       tracks,
		ReferenceOps:   references,
	}
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (s *Service) GetArtistsCascading(artistID ...uuid.UUID) ([]*models.ArtistEntry, error) {
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
