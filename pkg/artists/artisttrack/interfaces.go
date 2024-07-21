package artisttrack

// DatabaseOperations represents operations related to the database.
type DatabaseOperations interface {
	LinkArtistToTrack(artistId, trackId string)
	GetArtistToTrackLink(artistID string) (*ArtistTrack, error)
	GetArtistToTrackLinks(artistID ...string) ([]*ArtistTrack, error)
}
