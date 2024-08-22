package models

// Artist represents an artist with entity references.
type Artist struct {
	*ArtistEntry
	Tracks     []*Track
	References []*Reference
}
