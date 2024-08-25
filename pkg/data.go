package pkg

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// DataOps represents all DataAPI data operations.
type DataOps interface {
	GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error)
	GetTracks(trackID ...uuid.UUID) ([]*models.Track, error)
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
	GetReferences(internalID ...uuid.UUID) ([]*models.Reference, error)
	GetSources(sourceID ...uint) ([]*models.Source, error)
}

// DataAPI represents the main entry point to interact with functionalities for the defined entities.
type DataAPI struct {
	*SvcManager
}

// NewDataAPI creates a new instance of the DataAPI.
func NewDataAPI(database logic.DbOperations) DataOps {
	return &DataAPI{SvcManager: NewSvcManager(database)}
}

// GetArtists retrieves artists, with entity references, for the provided IDs.
func (d *DataAPI) GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error) {
	return d.GetArtistsCascading(artistID...)
}

// GetTracks retrieves tracks, with entity references, for the provided IDs.
func (d *DataAPI) GetTracks(trackID ...uuid.UUID) ([]*models.Track, error) {
	artistTrackEntries, err := d.GetArtistToTrackEntries(trackID...)
	if err != nil {
		return nil, err
	}

	artistIDs, _ := d.ExtractIDsFromArtistTrackEntries(artistTrackEntries)
	artists, err := d.GetArtistsCascading(artistIDs...)
	if err != nil {
		return nil, err
	}

	tracks := d.CollectTracks(artists)

	return tracks, nil
}
