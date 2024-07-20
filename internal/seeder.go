package internal

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"log"
)

// Seeder helps with deleting data from the database
type Seeder struct {
	Seeding *SeedingConfig
	API     *api.API
	Factory *DummyFactory
}

// NewSeeder creates a new instance of Seeder
func NewSeeder(seeding *SeedingConfig, api *api.API) *Seeder {
	return &Seeder{
		Seeding: seeding,
		API:     api,
		Factory: NewFactory(seeding.Sources, seeding.Instruments, seeding.Difficulties)}
}

// Seed attempts to seed the database with the minimally required values and dummy data.
func (s *Seeder) Seed() {
	s.minimumSeed()
	s.dummySeed()
}

// minimumSeed when enabled, seeds the database with the minimally required values.
func (s *Seeder) minimumSeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableMinimumSeedingCommand) {
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
func (s *Seeder) seedAdmin() {
	email := logic.GetEnvValueAsString(KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(KeyAdminInitialPassword)
	if email != "" && password != "" {
		s.API.UsersAPI().InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// seedInstruments seeds the instruments table with the default instruments.
func (s *Seeder) seedInstruments() {
	s.API.InstrumentsAPI().InsertInstruments(s.Seeding.Instruments...)
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *Seeder) seedDifficulties() {
	s.API.DifficultiesAPI().InsertDifficulties(s.Seeding.Difficulties...)
}

// seedSources seeds the sources from the config file.
func (s *Seeder) seedSources() {
	s.API.SourcesAPI().InsertSources(s.Seeding.Sources...)
}

// seedEndpoints seeds the endpoints from the config file.
func (s *Seeder) seedEndpoints() {
	s.API.EndpointsAPI().InsertEndpoints(s.Seeding.Endpoints...)
}

// dummySeed when enabled, seeds the database with dummy data.
func (s *Seeder) dummySeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableDummySeedingCommand) {
		log.Println("Did not seed the database with dummy data, as it was disabled.")
		return
	}
	artists := s.Factory.CreateArtists(s.Seeding.Dummies.Artists)
	s.insertArtists(artists)
}

// insertArtists inserts the given artists and references into the database.
func (s *Seeder) insertArtists(artists []*art.Artist) {
	for _, artist := range artists {
		s.API.ArtistsAPI().InsertArtist(artist)
		artistRef := s.Factory.CreateReferenceID(artist.ID, CategoryMusic, CategoryArtist)
		s.API.ReferencesAPI().InsertReference(artistRef)
		s.insertTracks(artist.Tracks, artist.ID)
	}
}

// insertTracks inserts the given tracks and references into the database.
func (s *Seeder) insertTracks(tracks []*trk.Track, artistID uuid.UUID) {
	for _, track := range tracks {
		s.API.TracksAPI().InsertTrack(track)
		s.API.ArtistTrackAPI().LinkArtistToTrack(artistID.String(), track.ID.String())
		trackRef := s.Factory.CreateReferenceID(track.ID, CategoryMusic, CategoryTrack)
		s.API.ReferencesAPI().InsertReference(trackRef)
		s.insertTabs(track.Tabs, track.ID)
	}
}

// insertTabs inserts the given tabs and references into the database.
func (s *Seeder) insertTabs(tabs []*tabs.Tab, trackID uuid.UUID) {
	for _, tab := range tabs {
		s.API.TabsAPI().InsertTab(tab)
		s.API.TrackTabAPI().LinkTrackToTab(trackID.String(), tab.ID.String())
		tabRef := s.Factory.CreateReferenceID(tab.ID, CategoryTabs, CategoryTab)
		s.API.ReferencesAPI().InsertReference(tabRef)
	}
}
