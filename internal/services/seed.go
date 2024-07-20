package services

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"log"
)

// SeedService helps with seeding data into the database
type SeedService struct {
	Seeding *internal.SeedingConfig
	API     *api.API
	Dummy   *DummyService
}

// NewSeedService creates a new instance of SeedService
func NewSeedService(seeding *internal.SeedingConfig, api *api.API) *SeedService {
	return &SeedService{
		Seeding: seeding,
		API:     api,
		Dummy:   NewDummyService(seeding.Sources, seeding.Instruments, seeding.Difficulties)}
}

// Seed attempts to seed the database with the minimally required values and dummy data.
func (s *SeedService) Seed() {
	s.minimumSeed()
	s.dummySeed()
}

// minimumSeed when enabled, seeds the database with the minimally required values.
func (s *SeedService) minimumSeed() {
	if !logic.GetEnvValueAsBoolean(internal.KeyDatabaseEnableMinimumSeedingCommand) {
		log.Println("Did not seed the database with the minimally required values, as it was disabled.")
		return
	}
	s.seedAdmin()
	s.seedInstruments()
	s.seedDifficulties()
	s.seedSources()
	s.seedEndpoints()
}

// seedAdmin inserts an administrator user into the database.
func (s *SeedService) seedAdmin() {
	email := logic.GetEnvValueAsString(internal.KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(internal.KeyAdminInitialPassword)
	if email != "" && password != "" {
		s.API.UsersAPI().InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// seedInstruments seeds the instruments table with the default instruments.
func (s *SeedService) seedInstruments() {
	s.API.InstrumentsAPI().InsertInstruments(s.Seeding.Instruments...)
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedService) seedDifficulties() {
	s.API.DifficultiesAPI().InsertDifficulties(s.Seeding.Difficulties...)
}

// seedSources seeds the sources from the config file.
func (s *SeedService) seedSources() {
	s.API.SourcesAPI().InsertSources(s.Seeding.Sources...)
}

// seedEndpoints seeds the endpoints from the config file.
func (s *SeedService) seedEndpoints() {
	s.API.EndpointsAPI().InsertEndpoints(s.Seeding.Endpoints...)
}

// dummySeed when enabled, seeds the database with dummy data.
func (s *SeedService) dummySeed() {
	if !logic.GetEnvValueAsBoolean(internal.KeyDatabaseEnableDummySeedingCommand) {
		log.Println("Did not seed the database with dummy data, as it was disabled.")
		return
	}
	artists := s.Dummy.CreateArtists(s.Seeding.Dummies.Artists)
	s.insertArtists(artists)
}

// insertArtists inserts the given artists and references into the database.
func (s *SeedService) insertArtists(artists []*art.Artist) {
	for _, artist := range artists {
		s.API.ArtistsAPI().InsertArtist(artist)
		artistRef := s.Dummy.CreateReferenceID(artist.ID, internal.CategoryMusic, internal.CategoryArtist)
		s.API.ReferencesAPI().InsertReference(artistRef)
		s.insertTracks(artist.Tracks, artist.ID)
	}
}

// insertTracks inserts the given tracks and references into the database.
func (s *SeedService) insertTracks(tracks []*trk.Track, artistID uuid.UUID) {
	for _, track := range tracks {
		s.API.TracksAPI().InsertTrack(track)
		s.API.ArtistTrackAPI().LinkArtistToTrack(artistID.String(), track.ID.String())
		trackRef := s.Dummy.CreateReferenceID(track.ID, internal.CategoryMusic, internal.CategoryTrack)
		s.API.ReferencesAPI().InsertReference(trackRef)
		s.insertTabs(track.Tabs, track.ID)
	}
}

// insertTabs inserts the given tabs and references into the database.
func (s *SeedService) insertTabs(tabs []*tabs.Tab, trackID uuid.UUID) {
	for _, tab := range tabs {
		s.API.TabsAPI().InsertTab(tab)
		s.API.TrackTabAPI().LinkTrackToTab(trackID.String(), tab.ID.String())
		tabRef := s.Dummy.CreateReferenceID(tab.ID, internal.CategoryTabs, internal.CategoryTab)
		s.API.ReferencesAPI().InsertReference(tabRef)
	}
}
