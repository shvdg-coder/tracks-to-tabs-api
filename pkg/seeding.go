package pkg

import (
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"log"
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
	artists := s.SeedArtists()
	tracks := s.SeedTracks(artists)
	s.SeedTabs(tracks)
}

// SeedInstruments seeds the instruments table with the default instruments.
func (s *SeedingAPI) SeedInstruments() {
	err := s.InsertInstrumentEntries(s.SeedingConfig.Instruments...)
	if err != nil {
		log.Fatalf("Failed to insert instruments: %s", err.Error())
	}
}

// SeedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedingAPI) SeedDifficulties() {
	err := s.InsertDifficultyEntries(s.SeedingConfig.Difficulties...)
	if err != nil {
		log.Fatalf("Failed to insert difficulties: %s", err.Error())
	}
}

// SeedSources seeds the sources from the config file.
func (s *SeedingAPI) SeedSources() {
	err := s.InsertSourceEntries(s.SeedingConfig.Sources...)
	if err != nil {
		log.Fatalf("Failed to insert sources: %s", err.Error())
	}
}

// SeedEndpoints seeds the endpoints from the config file.
func (s *SeedingAPI) SeedEndpoints() {
	err := s.InsertEndpointEntries(s.SeedingConfig.Endpoints...)
	if err != nil {
		log.Fatalf("Failed to insert endpoints: %s", err.Error())
	}
}

// SeedArtists seeds the artists according to the dummy settings in the config file and returns their IDs.
func (s *SeedingAPI) SeedArtists() []*models.ArtistEntry {
	dummyArtists := s.CreateArtists(s.Dummies.Artists)
	artistRefs := make([]*models.ReferenceEntry, 0)

	for _, artist := range dummyArtists {
		sourceMusic := s.GetRandomSource(CategoryMusic)
		artistIDRef := s.CreateReference(artist.ID, sourceMusic.ID, TypeID, CategoryArtist, s.CreateRandomUUID())
		artistRefs = append(artistRefs, artistIDRef)

		sourceTabs := s.GetRandomSource(CategoryTabs)
		artistNameRef := s.CreateReference(artist.ID, sourceTabs.ID, TypeName, CategoryArtist, s.formatName(artist.Name))
		artistRefs = append(artistRefs, artistNameRef)
	}

	err := s.InsertArtistEntries(dummyArtists...)
	if err != nil {
		log.Fatalf("Failed to insert artist references: %s", err.Error())
	}

	err = s.InsertReferenceEntries(artistRefs...)
	if err != nil {
		log.Fatalf("Failed to insert artist references: %s", err.Error())
	}

	return dummyArtists
}

// SeedTracks seeds the tracks according to the dummy settings in the config file and returns their IDs.
func (s *SeedingAPI) SeedTracks(artists []*models.ArtistEntry) []*models.TrackEntry {
	dummyTracks := make([]*models.TrackEntry, 0)
	dummyArtistTracks := make([]*models.ArtistTrackEntry, 0)
	trackRefs := make([]*models.ReferenceEntry, 0)

	for _, artist := range artists {
		tracks := s.CreateTracks(s.Dummies.Artists.Tracks)
		dummyTracks = append(dummyTracks, tracks...)

		artistTracks := s.CreateArtistTrackEntries(artist, tracks)
		dummyArtistTracks = append(dummyArtistTracks, artistTracks...)
	}

	for _, track := range dummyTracks {
		sourceMusic := s.GetRandomSource(CategoryMusic)
		trackIDRef := s.CreateReference(track.ID, sourceMusic.ID, TypeID, CategoryTrack, s.CreateRandomUUID())
		trackRefs = append(trackRefs, trackIDRef)

		sourceTabs := s.GetRandomSource(CategoryTabs)
		trackNameRef := s.CreateReference(track.ID, sourceTabs.ID, TypeName, CategoryTrack, s.formatName(track.Title))
		trackRefs = append(trackRefs, trackNameRef)
	}

	err := s.InsertTrackEntries(dummyTracks...)
	if err != nil {
		log.Fatalf("Failed to insert tracks: %s", err.Error())
	}

	err = s.InsertArtistTrackEntries(dummyArtistTracks...)
	if err != nil {
		log.Fatalf("Failed to insert artist tracks: %s", err.Error())
	}

	err = s.InsertReferenceEntries(trackRefs...)
	if err != nil {
		log.Fatalf("Failed to insert track references: %s", err.Error())
	}

	return dummyTracks
}

// SeedTabs seeds the tabs according to the dummy settings in the config file.
func (s *SeedingAPI) SeedTabs(tracks []*models.TrackEntry) []*models.TabEntry {
	dummyTabs := make([]*models.TabEntry, 0)
	dummyTrackTabs := make([]*models.TrackTabEntry, 0)
	tabRefs := make([]*models.ReferenceEntry, 0)

	for _, track := range tracks {
		tabs := s.CreateTabs(s.Dummies.Artists.Tracks.Tabs)
		dummyTabs = append(dummyTabs, tabs...)

		trackTabs := s.CreateTrackTabEntries(track, tabs)
		dummyTrackTabs = append(dummyTrackTabs, trackTabs...)
	}

	for _, tab := range dummyTabs {
		sourceTabs := s.GetRandomSource(CategoryTabs)

		tabIDRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeID, CategoryTab, s.CreateRandomUUID())
		tabRefs = append(tabRefs, tabIDRef)

		tabNameRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeName, CategoryTab, s.formatName(tab.Description))
		tabRefs = append(tabRefs, tabNameRef)
	}

	err := s.InsertTabEntries(dummyTabs...)
	if err != nil {
		log.Fatalf("Failed to insert tab entries: %v", err)
	}

	err = s.InsertTrackTabEntries(dummyTrackTabs...)
	if err != nil {
		log.Fatalf("Failed to insert track tabs: %s", err.Error())
	}

	err = s.InsertReferenceEntries(tabRefs...)
	if err != nil {
		log.Fatalf("Failed to insert tab references: %s", err.Error())
	}

	return dummyTabs
}

// formatName formats the provided name.
func (s *SeedingAPI) formatName(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
