package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/schemas"
)

// ArtistOps represents all operations related to artists.
type ArtistOps interface {
	schemas.ArtistSchema
	data.ArtistData
	mappers.ArtistMapper
	GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error)
	GetArtistsCascading(artistID ...uuid.UUID) ([]*models.Artist, error)
	ExtractIDsFromArtists(artists []*models.Artist) []uuid.UUID
	CollectTracks(artists []*models.Artist) []*models.Track
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

// GetArtists retrieves artists, without entity references, for the provided IDs.
func (a *ArtistSvc) GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error) {
	artistEntries, err := a.GetArtistsEntries(artistID...)
	if err != nil {
		return nil, err
	}

	artists := a.ArtistEntriesToArtists(artistEntries)

	return artists, nil
}

// GetArtistsCascading retrieves artists, with entity references, for the provided IDs.
func (a *ArtistSvc) GetArtistsCascading(artistID ...uuid.UUID) ([]*models.Artist, error) {
	artists, err := a.GetArtists(artistID...)
	if err != nil {
		return nil, err
	}

	err = a.LoadTracks(artists...)
	if err != nil {
		return nil, err
	}

	err = a.LoadReferences(artists...)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

// LoadTracks loads the models.Track's for the given models.Artist's.
func (a *ArtistSvc) LoadTracks(artists ...*models.Artist) error {
	artistIDs := a.ExtractIDsFromArtists(artists)
	artistTracksEntries, err := a.GetArtistToTrackEntries(artistIDs...)
	if err != nil {
		return err
	}

	_, trackIDs := a.ExtractIDsFromArtistTrackEntries(artistTracksEntries)
	tracks, err := a.GetTracksCascading(trackIDs...)
	if err != nil {
		return err
	}

	artistsMap := a.ArtistsToMap(artists)
	tracksMap := a.TracksToMap(tracks)

	a.MapArtistsToTracks(tracksMap, artistsMap, artistTracksEntries)
	a.MapTracksToArtists(artistsMap, tracksMap, artistTracksEntries)

	return nil
}

// LoadReferences loads references for the given artists.
func (a *ArtistSvc) LoadReferences(artists ...*models.Artist) error {
	artistIDs := a.ExtractIDsFromArtists(artists)
	references, err := a.GetReferencesCascading(artistIDs...)
	if err != nil {
		return err
	}

	artistsMap := a.ArtistsToMap(artists)
	a.MapReferencesToArtists(artistsMap, references)

	return nil
}

// ExtractIDsFromArtists retrieves the ID's from the models.Artist's.
func (a *ArtistSvc) ExtractIDsFromArtists(artists []*models.Artist) []uuid.UUID {
	artistIDs := make([]uuid.UUID, len(artists))
	for i, artist := range artists {
		artistIDs[i] = artist.ID
	}
	return artistIDs
}

// CollectTracks plucks the models.Track's from each of the models.Artist's, and returns them.
func (a *ArtistSvc) CollectTracks(artists []*models.Artist) []*models.Track {
	tracks := make([]*models.Track, 0)
	for _, artist := range artists {
		tracks = append(tracks, artist.Tracks...)
	}
	return tracks
}
