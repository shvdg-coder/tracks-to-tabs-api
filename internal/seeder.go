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
	Seeding *Seeding
	API     *api.API
	Factory *DummyFactory
}

// NewSeeder creates a new instance of Seeder
func NewSeeder(seeding *Seeding, api *api.API) *Seeder {
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
		s.API.Users.InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// seedInstruments seeds the instruments table with the default instruments.
func (s *Seeder) seedInstruments() {
	s.API.Instruments.InsertInstruments(s.Seeding.Instruments...)
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *Seeder) seedDifficulties() {
	s.API.Difficulties.InsertDifficulties(s.Seeding.Difficulties...)
}

// seedSources seeds the sources from the config file.
func (s *Seeder) seedSources() {
	s.API.Sources.InsertSources(s.Seeding.Sources...)
}

// seedEndpoints seeds the endpoints from the config file.
func (s *Seeder) seedEndpoints() {
	s.API.Endpoints.InsertEndpoints(s.Seeding.Endpoints...)
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
		s.API.Artists.InsertArtist(artist)
		artistRef := s.Factory.CreateReferenceID(artist.ID, "artists")
		s.API.References.InsertReference(artistRef)
		s.insertTracks(artist.Tracks, artist.ID)
	}
}

// insertTracks inserts the given tracks and references into the database.
func (s *Seeder) insertTracks(tracks []*trk.Track, artistID uuid.UUID) {
	for _, track := range tracks {
		s.API.Tracks.InsertTrack(track)
		s.API.Artists.LinkArtistToTrack(artistID.String(), track.ID.String())
		trackRef := s.Factory.CreateReferenceID(track.ID, "tracks")
		s.API.References.InsertReference(trackRef)
		s.insertTabs(track.Tabs, track.ID)
	}
}

// insertTabs inserts the given tabs and references into the database.
func (s *Seeder) insertTabs(tabs []*tabs.Tab, trackID uuid.UUID) {
	for _, tab := range tabs {
		s.API.Tabs.InsertTab(tab)
		s.API.Tracks.LinkTrackToTab(trackID.String(), tab.ID.String())
		tabRef := s.Factory.CreateReferenceID(tab.ID, "tabs")
		s.API.References.InsertReference(tabRef)
	}
}

//TODO: Put error messages somewhere (constants), string format
//TODO: Better deal with errors and logging
//TODO: Add unit tests
//TODO: Add integration tests
//TODO: Create views?
