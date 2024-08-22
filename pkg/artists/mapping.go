package artists

import (
	"github.com/google/uuid"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artisttrack"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// MappingOperations represents operations related to artist data mapping.
type MappingOperations interface {
	ArtistsToMap(artists []*Artist) map[uuid.UUID]*Artist
	MapTracksToArtists(artistTracks []*arttrk.ArtistTrack, artistsMap map[uuid.UUID]*Artist, tracksMap map[uuid.UUID]*trk.Track) []*Artist
	MapReferencesToArtists(artistsMap map[uuid.UUID]*Artist, references []*ref.Reference) []*Artist
}

// MappingService is responsible for mapping entities to artists.
type MappingService struct {
	MappingOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService() MappingOperations {
	return &MappingService{}
}

// ArtistsToMap transforms a slice of artists into a map where the key is the ID and the value the Artist.
func (m *MappingService) ArtistsToMap(artists []*Artist) map[uuid.UUID]*Artist {
	artistMap := make(map[uuid.UUID]*Artist)
	for _, artist := range artists {
		artistMap[artist.ID] = artist
	}
	return artistMap
}

// MapTracksToArtists adds the tracks to the artist.
func (m *MappingService) MapTracksToArtists(artistTracks []*arttrk.ArtistTrack, artistsMap map[uuid.UUID]*Artist, tracksMap map[uuid.UUID]*trk.Track) []*Artist {
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
	var artists []*Artist
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}

// MapReferencesToArtists maps references.Reference's to Artist's.
func (m *MappingService) MapReferencesToArtists(artistsMap map[uuid.UUID]*Artist, references []*ref.Reference) []*Artist {
	for _, reference := range references {
		artist, ok := artistsMap[reference.InternalID]
		if !ok {
			continue
		}
		artist.References = append(artist.References, reference)
	}
	var artists []*Artist
	for _, artist := range artistsMap {
		artists = append(artists, artist)
	}
	return artists
}
