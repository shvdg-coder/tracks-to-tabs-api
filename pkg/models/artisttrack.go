package models

// ArtistTrack represents an 'artist track' link with entity references.
type ArtistTrack struct {
	Artist *Artist
	Track  *Track
}
