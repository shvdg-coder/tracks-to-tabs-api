package pkg

import (
	"github.com/google/uuid"
	"strings"
)

// SeedOps represents all operations related to seeding.
type SeedOps interface {
	Seed()
}

// SeedingAPI is responsible for seeding data.
type SeedingAPI struct {
	SvcOps
	*SeedingConfig
	DummyOps
}

// NewSeedingAPI instantiates a SeedingAPI.
func NewSeedingAPI(svcManager SvcOps, config *SeedingConfig, dummies DummyOps) SeedOps {
	return &SeedingAPI{svcManager, config, dummies}
}

// Seed seeds the database with the entries found in the provided SeedingConfig.
func (s *SeedingAPI) Seed() {
	s.SeedInstruments()
	s.SeedDifficulties()
	s.SeedSources()
	s.SeedEndpoints()
	artistIDs := s.SeedArtists()
	trackIDs := s.SeedTracks(artistIDs)
	s.SeedTabs(trackIDs)
}

// SeedInstruments seeds the instruments table with the default instruments.
func (s *SeedingAPI) SeedInstruments() {
	s.InsertInstrumentEntries(s.SeedingConfig.Instruments...)
}

// SeedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedingAPI) SeedDifficulties() {
	s.InsertDifficultyEntries(s.SeedingConfig.Difficulties...)
}

// SeedSources seeds the sources from the config file.
func (s *SeedingAPI) SeedSources() {
	s.InsertSourceEntries(s.SeedingConfig.Sources...)
}

// SeedEndpoints seeds the endpoints from the config file.
func (s *SeedingAPI) SeedEndpoints() {
	s.InsertEndpointEntries(s.SeedingConfig.Endpoints...)
}

// SeedArtists seeds the artists according to the dummy settings in the config file and returns their IDs.
func (s *SeedingAPI) SeedArtists() []uuid.UUID {
	var artistIDs []uuid.UUID
	return artistIDs
}

// SeedTracks seeds the tracks according to the dummy settings in the config file and returns their IDs.
func (s *SeedingAPI) SeedTracks(artistIDs []uuid.UUID) []uuid.UUID {
	var trackIDs []uuid.UUID

	for _, artistID := range artistIDs {
		dummyTracks := s.CreateTracks(s.Dummies.Artists.Tracks)
		for _, track := range dummyTracks {
			s.InsertTrackEntry(track)
			s.LinkArtistToTrack(artistID, track.ID)

			sourceMusic := s.GetRandomSource(CategoryMusic)
			trackIDRef := s.CreateReference(track.ID, sourceMusic.ID, TypeID, CategoryTrack, s.CreateRandomUUID())
			s.InsertReferenceEntry(trackIDRef)

			sourceTabs := s.GetRandomSource(CategoryTabs)
			trackNameRef := s.CreateReference(track.ID, sourceTabs.ID, TypeName, CategoryTrack, s.formatName(track.Title))
			s.InsertReferenceEntry(trackNameRef)

			trackIDs = append(trackIDs, track.ID)
		}
	}

	return trackIDs
}

// SeedTabs seeds the tabs according to the dummy settings in the config file.
func (s *SeedingAPI) SeedTabs(trackIDs []uuid.UUID) {
	for _, trackID := range trackIDs {
		dummyTabs := s.CreateTabs(s.Dummies.Artists.Tracks.Tabs)
		for _, tab := range dummyTabs {
			s.InsertTabEntry(tab)
			s.LinkTrackToTab(trackID, tab.ID)

			sourceTabs := s.GetRandomSource(CategoryTabs)

			tabIDRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeID, CategoryTab, s.CreateRandomUUID())
			s.InsertReferenceEntry(tabIDRef)

			tabNameRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeName, CategoryTab, s.formatName(tab.Description))
			s.InsertReferenceEntry(tabNameRef)
		}
	}
}

// formatName formats the provided name.
func (s *SeedingAPI) formatName(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
