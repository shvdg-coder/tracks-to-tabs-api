package internal

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"log"
)

// SeedingOperations represents operations related to seeding the database.
type SeedingOperations interface {
	SeedAll()
	SeedMinimumData()
	SeedUsers()
	SeedInstruments()
	SeedDifficulties()
	SeedSources()
	SeedEndpoints()
	SeedDummyData()
	SeedArtists(artists []*art.Artist)
	SeedTracks(tracks []*trk.Track, artistID uuid.UUID)
	SeedTabs(tabs []*tbs.Tab, trackID uuid.UUID)
}

// SeedService helps with seeding data into the database
type SeedService struct {
	Seeding *SeedingConfig
	API     *api.API
	Dummy   *DummyService
}

// NewSeedService creates a new instance of SeedService
func NewSeedService(seeding *SeedingConfig, api *api.API) *SeedService {
	return &SeedService{
		Seeding: seeding,
		API:     api,
		Dummy:   NewDummyService(seeding.Sources, seeding.Instruments, seeding.Difficulties)}
}

// SeedAll attempts to seed the database with the minimally required values and dummy data.
func (s *SeedService) SeedAll() {
	s.SeedMinimumData()
	s.SeedDummyData()
}

// SeedMinimumData when enabled, seeds the database with the minimally required values.
func (s *SeedService) SeedMinimumData() {
	s.SeedUsers()
	s.SeedInstruments()
	s.SeedDifficulties()
	s.SeedSources()
	s.SeedEndpoints()
}

// SeedUsers inserts an administrator user into the database.
func (s *SeedService) SeedUsers() {
	email := logic.GetEnvValueAsString(KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(KeyAdminInitialPassword)
	if email != "" && password != "" {
		s.API.Users().InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// SeedInstruments seeds the instruments table with the default instruments.
func (s *SeedService) SeedInstruments() {
	s.API.Instruments().InsertInstruments(s.Seeding.Instruments...)
}

// SeedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedService) SeedDifficulties() {
	s.API.Difficulties().InsertDifficulties(s.Seeding.Difficulties...)
}

// SeedSources seeds the sources from the config file.
func (s *SeedService) SeedSources() {
	s.API.Sources().InsertSources(s.Seeding.Sources...)
}

// SeedEndpoints seeds the endpoints from the config file.
func (s *SeedService) SeedEndpoints() {
	s.API.Endpoints().InsertEndpoints(s.Seeding.Endpoints...)
}

// SeedDummyData when enabled, seeds the database with dummy data.
func (s *SeedService) SeedDummyData() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableDummySeedingCommand) {
		log.Println("Did not seed the database with dummy data, as it was disabled.")
		return
	}
	artists := s.Dummy.CreateArtists(s.Seeding.Dummies.Artists)
	s.SeedArtists(artists)
}

// SeedArtists inserts the given artists and references into the database.
func (s *SeedService) SeedArtists(artists []*art.Artist) {
	for _, artist := range artists {
		s.API.Artists().InsertArtist(artist)
		artistRef := s.Dummy.CreateReferenceID(artist.ID, CategoryMusic, CategoryArtist)
		s.API.References().InsertReference(artistRef)
		s.SeedTracks(artist.Tracks, artist.ID)
	}
}

// SeedTracks inserts the given tracks and references into the database.
func (s *SeedService) SeedTracks(tracks []*trk.Track, artistID uuid.UUID) {
	for _, track := range tracks {
		s.API.Tracks().InsertTrack(track)
		s.API.Artists().LinkArtistToTrack(artistID, track.ID)
		trackRef := s.Dummy.CreateReferenceID(track.ID, CategoryMusic, CategoryTrack)
		s.API.References().InsertReference(trackRef)
		s.SeedTabs(track.Tabs, track.ID)
	}
}

// SeedTabs inserts the given tabs and references into the database.
func (s *SeedService) SeedTabs(tabs []*tbs.Tab, trackID uuid.UUID) {
	for _, tab := range tabs {
		s.API.Tabs().InsertTab(tab)
		s.API.Tracks().LinkTrackToTab(trackID, tab.ID)
		tabRef := s.Dummy.CreateReferenceID(tab.ID, CategoryTabs, CategoryTab)
		s.API.References().InsertReference(tabRef)
	}
}
