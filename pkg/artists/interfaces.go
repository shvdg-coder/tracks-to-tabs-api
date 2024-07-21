package artists

import arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"

// DatabaseOperations represents operations related to the database.
type DatabaseOperations interface {
	InsertArtist(artist *Artist)
	InsertArtists(artist ...*Artist)
	GetArtist(artistID string) (*Artist, error)
	GetArtists(artistID ...string) ([]*Artist, error)
}

// MappingOperations represents operations related to data mapping.
type MappingOperations interface {
	GetArtistsCascading(artistID ...string) ([]*Artist, error)
}

// Operations represents all operations related to Artists.
type Operations interface {
	DatabaseOperations
	MappingOperations
	arttrk.DatabaseOperations
}
