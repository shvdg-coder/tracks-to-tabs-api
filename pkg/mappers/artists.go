package mappers

import (
	"github.com/google/uuid"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistMapper represents operations related to artist data mapping.
type ArtistMapper interface {
	ArtistsToMap(artists []*trk.Artist) map[uuid.UUID]*trk.Artist
	MapTracksToArtists(artistTracks []*trk.ArtistTrackEntry, artistsMap map[uuid.UUID]*trk.Artist, tracksMap map[uuid.UUID]*trk.Track) map[uuid.UUID]*trk.Artist
	MapReferencesToArtists(artistsMap map[uuid.UUID]*trk.Artist, references []*trk.Reference) map[uuid.UUID]*trk.Artist
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
func (m *ArtistSvc) MapTracksToArtists(artistTracks []*trk.ArtistTrackEntry, artistsMap map[uuid.UUID]*trk.Artist, tracksMap map[uuid.UUID]*trk.Track) map[uuid.UUID]*trk.Artist {
	for _, artistTrack := range artistTracks {
		artist := artistsMap[artistTrack.ArtistID]
		track := tracksMap[artistTrack.TrackID]
		artist.Tracks = append(artist.Tracks, track)
	}
	return artistsMap
}

// MapReferencesToArtists maps references.Reference's to ArtistEntry's.
func (m *ArtistSvc) MapReferencesToArtists(artistsMap map[uuid.UUID]*trk.Artist, references []*trk.Reference) map[uuid.UUID]*trk.Artist {
	for _, reference := range references {
		artist := artistsMap[reference.InternalID]
		artist.References = append(artist.References, reference)
	}
	return artistsMap
}
