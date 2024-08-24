package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// ArtistOps represents all operations related to artists.
type ArtistOps interface {
	schemas.ArtistSchema
	data.ArtistData
	mappers.ArtistMapper
	GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error)
}

// ArtistSvc is responsible for managing and retrieving artists.
type ArtistSvc struct {
	schemas.ArtistSchema
	data.ArtistData
	mappers.ArtistMapper
	ArtistTrackOps
	TrackOps
	ReferenceOps
}

// NewArtistSvc instantiates a ArtistSvc.
func NewArtistSvc(schema schemas.ArtistSchema, data data.ArtistData, mapper mappers.ArtistMapper, artistTracks ArtistTrackOps, tracks TrackOps, references ReferenceOps) ArtistOps {
	return &ArtistSvc{
		ArtistSchema:   schema,
		ArtistData:     data,
		ArtistMapper:   mapper,
		ArtistTrackOps: artistTracks,
		TrackOps:       tracks,
		ReferenceOps:   references,
	}
}

// GetArtists retrieves artists, with entity references, for the provided IDs.
func (a *ArtistSvc) GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error) {
	artistEntries, err := a.GetArtistsEntries(artistID...)
	if err != nil {
		return nil, err
	}

	artistTracksEntries, err := a.GetArtistToTrackEntries(artistID...)
	if err != nil {
		return nil, err
	}

	trackIDs := a.ExtractTrackIDs(artistTracksEntries)
	tracks, err := a.GetTracks(trackIDs...)
	if err != nil {
		return nil, err
	}

	references, err := a.GetReferences(artistID...)
	if err != nil {
		return nil, err
	}

	artists := a.ArtistEntriesToArtists(artistEntries)
	artistsMap := a.ArtistsToMap(artists)
	tracksMap := a.TracksToMap(tracks)

	artistsMap = a.MapTracksToArtists(artistsMap, tracksMap, artistTracksEntries)
	artistsMap = a.MapReferencesToArtists(artistsMap, references)
	artists = a.MapToArtists(artistsMap)

	return artists, nil
}
