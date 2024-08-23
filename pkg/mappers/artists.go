package mappers

import (
	"github.com/google/uuid"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistMapper represents operations related to artist data mapping.
type ArtistMapper interface {
	ArtistsToMap(artists []*trk.Artist) map[uuid.UUID]*trk.Artist
	MapTracksToArtists(artistTracks []*trk.ArtistTrackEntry, artistsMap map[uuid.UUID]*trk.Artist, tracksMap map[uuid.UUID]*trk.Track) []*trk.Artist
	MapReferencesToArtists(artistsMap map[uuid.UUID]*trk.Artist, references []*trk.Reference) []*trk.Artist
}

// ArtistSvc is responsible for mapping entities to artists.
type ArtistSvc struct {
	ArtistMapper
}

// NewArtistSvc creates a new instance of ArtistSvc.
func NewArtistSvc() ArtistMapper {
	return &ArtistSvc{}
}

// ArtistsToMap transforms a slice of artists into a map where the key is the ID and the value the ArtistEntry.
func (m *ArtistSvc) ArtistsToMap(artists []*trk.Artist) map[uuid.UUID]*trk.Artist {
	artistMap := make(map[uuid.UUID]*trk.Artist)
	for _, artist := range artists {
		artistMap[artist.ID] = artist
	}
	return artistMap
}

// MapTracksToArtists adds the tracks to the artist.
func (m *ArtistSvc) MapTracksToArtists(artistTracks []*trk.ArtistTrackEntry, artistsMap map[uuid.UUID]*trk.Artist, tracksMap map[uuid.UUID]*trk.Track) []*trk.Artist {
	for _, link := range artistTracks {
		artist := artistsMap[link.ArtistID]
		track := tracksMap[link.TrackID]
		artist.Tracks = append(artist.Tracks, track)
	}
	var artists []*trk.Artist
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}

// MapReferencesToArtists maps references.Reference's to ArtistEntry's.
func (m *ArtistSvc) MapReferencesToArtists(artistsMap map[uuid.UUID]*trk.Artist, references []*trk.Reference) []*trk.Artist {
	for _, reference := range references {
		artist := artistsMap[reference.InternalID]
		artist.References = append(artist.References, reference)
	}
	var artists []*trk.Artist
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}
