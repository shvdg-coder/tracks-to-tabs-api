package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
)

// DataOps represents all DataSvc data operations.
type DataOps interface {
	GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error)
	GetTracks(trackID ...uuid.UUID) ([]*models.Track, error)
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
	SvcOps
}

// DataSvc represents the main entry point to interact with functionalities for the defined entities.
type DataSvc struct {
	SvcOps
}

// NewDataSvc creates a new instance of the DataSvc.
func NewDataSvc(svcManager SvcOps) DataOps {
	return &DataSvc{svcManager}
}

// GetArtists retrieves artists, with entity references, for the provided IDs.
func (d *DataSvc) GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error) {
	artists, err := d.GetArtistsCascading(artistID...)
	if err != nil {
		return nil, err
	}

	d.LoadArtistsResources(artists...)

	return artists, nil
}

// GetTracks retrieves tracks, with entity references, for the provided IDs.
func (d *DataSvc) GetTracks(trackID ...uuid.UUID) ([]*models.Track, error) {
	artistTrackEntries, err := d.GetArtistToTrackEntries(trackID...)
	if err != nil {
		return nil, err
	}

	artistIDs, _ := d.ExtractIDsFromArtistTrackEntries(artistTrackEntries)
	artists, err := d.GetArtists(artistIDs...)
	if err != nil {
		return nil, err
	}

	tracks := d.CollectTracks(artists)
	d.LoadTracksResources(tracks...)

	return tracks, nil
}

// GetTabs retrieves tabs, with entity references, for the provided IDs.
func (d *DataSvc) GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error) {
	tabTrackEntries, err := d.GetTrackToTabEntries(tabID...)
	if err != nil {
		return nil, err
	}

	trackIDs, _ := d.ExtractIDsFromTrackTabEntries(tabTrackEntries)
	tracks, err := d.GetTracks(trackIDs...)
	if err != nil {
		return nil, err
	}

	tabs := d.CollectTabs(tracks)
	d.LoadTabsResources(tabs...)

	return tabs, nil
}
