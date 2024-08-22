package mappers

import (
	"github.com/google/uuid"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistOps represents operations related to artist data mapping.
type ArtistOps interface {
	ArtistsToMap(artists []*trk.ArtistEntry) map[uuid.UUID]*trk.ArtistEntry
	MapTracksToArtists(artistTracks []*trk.ArtistTrackEntry, artistsMap map[uuid.UUID]*trk.ArtistEntry, tracksMap map[uuid.UUID]*trk.Track) []*trk.ArtistEntry
	MapReferencesToArtists(artistsMap map[uuid.UUID]*trk.ArtistEntry, references []*trk.Reference) []*trk.ArtistEntry
}

// ArtistServ is responsible for mapping entities to artists.
type ArtistServ struct {
	ArtistOps
}

// NewArtistServ creates a new instance of ArtistServ.
func NewArtistServ() ArtistOps {
	return &ArtistServ{}
}

// ArtistsToMap transforms a slice of artists into a map where the key is the ID and the value the ArtistEntry.
func (m *ArtistServ) ArtistsToMap(artists []*trk.ArtistEntry) map[uuid.UUID]*trk.ArtistEntry {
	artistMap := make(map[uuid.UUID]*trk.ArtistEntry)
	for _, artist := range artists {
		artistMap[artist.ID] = artist
	}
	return artistMap
}

// MapTracksToArtists adds the tracks to the artist.
func (m *ArtistServ) MapTracksToArtists(artistTracks []*trk.ArtistTrackEntry, artistsMap map[uuid.UUID]*trk.ArtistEntry, tracksMap map[uuid.UUID]*trk.Track) []*trk.ArtistEntry {
	for _, link := range artistTracks {
		artist, ok := artistsMap[link.ArtistID]
		if !ok {
			continue
		}
		track, ok := tracksMap[link.TrackID]
		if !ok {
			continue
		}
		artist.Tracks = append(artist.Tracks, track)
	}
	var artists []*trk.ArtistEntry
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}

// MapReferencesToArtists maps references.Reference's to ArtistEntry's.
func (m *ArtistServ) MapReferencesToArtists(artistsMap map[uuid.UUID]*trk.ArtistEntry, references []*trk.Reference) []*trk.ArtistEntry {
	for _, reference := range references {
		artist, ok := artistsMap[reference.InternalID]
		if !ok {
			continue
		}
		artist.References = append(artist.References, reference)
	}
	var artists []*trk.ArtistEntry
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}
