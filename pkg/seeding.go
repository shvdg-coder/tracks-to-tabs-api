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

func (s *SeedingAPI) Seed() {
	s.SeedInstruments()
	s.SeedDifficulties()
	s.SeedSources()
	s.SeedEndpoints()
	artists := s.CreateAndInsertArtists()
	s.CreateAndInsertArtistReferences(artists)
	tracks := s.CreateAndInsertTracks(artists)
	s.CreateAndInsertTrackReferences(tracks)
	tabs := s.CreateAndInsertTabs(tracks)
	s.CreateAndInsertTabReferences(tabs)
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

// CreateAndInsertArtists creates and inserts artist entries and returns it.
func (s *SeedingAPI) CreateAndInsertArtists() []*models.ArtistEntry {
	dummyArtists := s.CreateArtists(s.Dummies.Artists)

	err := s.InsertArtistEntries(dummyArtists...)
	if err != nil {
		log.Fatalf("Failed to insert artist entries: %s", err.Error())
	}

	return dummyArtists
}

// CreateArtistReferences creates artist references and returns it.
func (s *SeedingAPI) CreateArtistReferences(artist *models.ArtistEntry) []*models.ReferenceEntry {
	artistRefs := make([]*models.ReferenceEntry, 0)

	sourceMusic := s.GetRandomSource(CategoryMusic)
	artistIDRef := s.CreateReference(artist.ID, sourceMusic.ID, TypeID, CategoryArtist, s.CreateRandomUUID())
	artistRefs = append(artistRefs, artistIDRef)

	sourceTabs := s.GetRandomSource(CategoryTabs)
	artistNameRef := s.CreateReference(artist.ID, sourceTabs.ID, TypeName, CategoryArtist, s.formatName(artist.Name))
	artistRefs = append(artistRefs, artistNameRef)

	return artistRefs
}

// CreateAndInsertArtistReferences creates and inserts artist references for given artists.
func (s *SeedingAPI) CreateAndInsertArtistReferences(artists []*models.ArtistEntry) {
	artistRefs := make([]*models.ReferenceEntry, 0)

	for _, artist := range artists {
		artistRefs = append(artistRefs, s.CreateArtistReferences(artist)...)
	}

	err := s.InsertReferenceEntries(artistRefs...)
	if err != nil {
		log.Fatalf("Failed to insert artist references: %s", err.Error())
	}
}

// CreateAndInsertTracks creates and inserts tracks for given artists and returns the tracks.
func (s *SeedingAPI) CreateAndInsertTracks(artists []*models.ArtistEntry) []*models.TrackEntry {
	dummyTracks := make([]*models.TrackEntry, 0)
	dummyArtistTracks := make([]*models.ArtistTrackEntry, 0)

	for _, artist := range artists {
		tracks := s.CreateTracks(s.Dummies.Artists.Tracks)
		dummyTracks = append(dummyTracks, tracks...)

		artistTracks := s.CreateArtistTrackEntries(artist, tracks)
		dummyArtistTracks = append(dummyArtistTracks, artistTracks...)
	}

	err := s.InsertTrackEntries(dummyTracks...)
	if err != nil {
		log.Fatalf("Failed to insert tracks: %s", err.Error())
	}

	err = s.InsertArtistTrackEntries(dummyArtistTracks...)
	if err != nil {
		log.Fatalf("Failed to insert artist tracks: %s", err.Error())
	}

	return dummyTracks
}

// CreateTrackReferences creates track references for a given track and returns it.
func (s *SeedingAPI) CreateTrackReferences(track *models.TrackEntry) []*models.ReferenceEntry {
	trackRefs := make([]*models.ReferenceEntry, 0)

	sourceMusic := s.GetRandomSource(CategoryMusic)
	trackIDRef := s.CreateReference(track.ID, sourceMusic.ID, TypeID, CategoryTrack, s.CreateRandomUUID())
	trackRefs = append(trackRefs, trackIDRef)

	sourceTabs := s.GetRandomSource(CategoryTabs)
	trackNameRef := s.CreateReference(track.ID, sourceTabs.ID, TypeName, CategoryTrack, s.formatName(track.Title))
	trackRefs = append(trackRefs, trackNameRef)

	return trackRefs
}

// CreateAndInsertTrackReferences creates and inserts track references for given tracks.
func (s *SeedingAPI) CreateAndInsertTrackReferences(tracks []*models.TrackEntry) {
	trackRefs := make([]*models.ReferenceEntry, 0)

	for _, track := range tracks {
		trackRefs = append(trackRefs, s.CreateTrackReferences(track)...)
	}

	err := s.InsertReferenceEntries(trackRefs...)
	if err != nil {
		log.Fatalf("Failed to insert track references: %s", err.Error())
	}
}

// CreateAndInsertTabs creates and inserts tabs for given tacks and returns the tabs.
func (s *SeedingAPI) CreateAndInsertTabs(tracks []*models.TrackEntry) []*models.TabEntry {
	dummyTabs := make([]*models.TabEntry, 0)
	dummyTrackTabs := make([]*models.TrackTabEntry, 0)

	for _, track := range tracks {
		tabs := s.CreateTabs(s.Dummies.Artists.Tracks.Tabs)
		dummyTabs = append(dummyTabs, tabs...)

		trackTabs := s.CreateTrackTabEntries(track, tabs)
		dummyTrackTabs = append(dummyTrackTabs, trackTabs...)
	}

	err := s.InsertTabEntries(dummyTabs...)
	if err != nil {
		log.Fatalf("Failed to insert tab entries: %v", err)
	}

	err = s.InsertTrackTabEntries(dummyTrackTabs...)
	if err != nil {
		log.Fatalf("Failed to insert track tabs: %s", err.Error())
	}

	return dummyTabs
}

// CreateTabReferences creates tab references for a given tab and returns it.
func (s *SeedingAPI) CreateTabReferences(tab *models.TabEntry) []*models.ReferenceEntry {
	tabRefs := make([]*models.ReferenceEntry, 0)

	sourceTabs := s.GetRandomSource(CategoryTabs)

	tabIDRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeID, CategoryTab, s.CreateRandomUUID())
	tabRefs = append(tabRefs, tabIDRef)

	tabNameRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeName, CategoryTab, s.formatName(tab.Description))
	tabRefs = append(tabRefs, tabNameRef)

	return tabRefs
}

// CreateAndInsertTabReferences creates and inserts tab references for given tabs.
func (s *SeedingAPI) CreateAndInsertTabReferences(tabs []*models.TabEntry) {
	tabRefs := make([]*models.ReferenceEntry, 0)

	for _, tab := range tabs {
		tabRefs = append(tabRefs, s.CreateTabReferences(tab)...)
	}

	err := s.InsertReferenceEntries(tabRefs...)
	if err != nil {
		log.Fatalf("Failed to insert tab references: %s", err.Error())
	}
}

// formatName formats the provided name.
func (s *SeedingAPI) formatName(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
