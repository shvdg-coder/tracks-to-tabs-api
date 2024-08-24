package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistMapper represents operations related to artist data mapping.
type ArtistMapper interface {
	ArtistEntriesToArtists(artistEntries []*models.ArtistEntry) []*models.Artist
	ArtistsToMap(artists []*models.Artist) map[uuid.UUID]*models.Artist
	MapToArtists(artistsMap map[uuid.UUID]*models.Artist) []*models.Artist
	MapTracksToArtists(artistsMap map[uuid.UUID]*models.Artist, tracksMap map[uuid.UUID]*models.Track, artistTracks []*models.ArtistTrackEntry) map[uuid.UUID]*models.Artist
	MapReferencesToArtists(artistsMap map[uuid.UUID]*models.Artist, references []*models.Reference) map[uuid.UUID]*models.Artist
}

// ArtistSvc is responsible for mapping entities to artists.
type ArtistSvc struct {
	ArtistMapper
}

// NewArtistSvc creates a new instance of ArtistSvc.
func NewArtistSvc() ArtistMapper {
	return &ArtistSvc{}
}

// ArtistEntriesToArtists transforms the models.ArtistEntry's to models.Artist's.
func (a *ArtistSvc) ArtistEntriesToArtists(artistEntries []*models.ArtistEntry) []*models.Artist {
	artists := make([]*models.Artist, len(artistEntries))
	for i, artistEntry := range artistEntries {
		artists[i] = &models.Artist{ArtistEntry: artistEntry}
	}
	return artists
}

// ArtistsToMap transforms a slice of models.Artist's into a map where the key is the ID and the value the models.Artist.
func (a *ArtistSvc) ArtistsToMap(artists []*models.Artist) map[uuid.UUID]*models.Artist {
	artistMap := make(map[uuid.UUID]*models.Artist, len(artists))
	for _, artist := range artists {
		artistMap[artist.ID] = artist
	}
	return artistMap
}

// MapToArtists transforms a map of models.Artist's into a slice of models.Artist's.
func (a *ArtistSvc) MapToArtists(artistsMap map[uuid.UUID]*models.Artist) []*models.Artist {
	artists := make([]*models.Artist, 0)
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}

// MapTracksToArtists adds the models.Track's to the models.Artist, by updating the provided models.Artist's map and returning it.
func (a *ArtistSvc) MapTracksToArtists(artistsMap map[uuid.UUID]*models.Artist, tracksMap map[uuid.UUID]*models.Track, artistTracks []*models.ArtistTrackEntry) map[uuid.UUID]*models.Artist {
	for _, artistTrack := range artistTracks {
		artist := artistsMap[artistTrack.ArtistID]
		track := tracksMap[artistTrack.TrackID]
		artist.Tracks = append(artist.Tracks, track)
	}
	return artistsMap
}

// MapReferencesToArtists maps models.Reference's to models.Artist's, by updating the provided models.Artist's map and returning it.
func (a *ArtistSvc) MapReferencesToArtists(artistsMap map[uuid.UUID]*models.Artist, references []*models.Reference) map[uuid.UUID]*models.Artist {
	for _, reference := range references {
		artist := artistsMap[reference.InternalID]
		artist.References = append(artist.References, reference)
	}
	return artistsMap
}
