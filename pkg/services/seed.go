package services

import (
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/constants"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"log"
	"strings"
)

// SeedOps represents all operations related to seeding.
type SeedOps interface {
	Seed()
}

// SeedSvc is responsible for seeding data.
type SeedSvc struct {
	SvcOps
	*models.SeedingConfig
	DummyOps
}

// NewSeedSvc instantiates a SeedSvc.
func NewSeedSvc(svcManager SvcOps, config *models.SeedingConfig, dummies DummyOps) SeedOps {
	return &SeedSvc{svcManager, config, dummies}
}

func (s *SeedSvc) Seed() {
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
func (s *SeedSvc) SeedInstruments() {
	err := s.InsertInstrumentEntries(s.SeedingConfig.Instruments...)
	if err != nil {
		log.Fatalf("Failed to insert instruments: %s", err.Error())
	}
}

// SeedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedSvc) SeedDifficulties() {
	err := s.InsertDifficultyEntries(s.SeedingConfig.Difficulties...)
	if err != nil {
		log.Fatalf("Failed to insert difficulties: %s", err.Error())
	}
}

// SeedSources seeds the sources from the config file.
func (s *SeedSvc) SeedSources() {
	err := s.InsertSourceEntries(s.SeedingConfig.Sources...)
	if err != nil {
		log.Fatalf("Failed to insert sources: %s", err.Error())
	}
}

// SeedEndpoints seeds the endpoints from the config file.
func (s *SeedSvc) SeedEndpoints() {
	err := s.InsertEndpointEntries(s.SeedingConfig.Endpoints...)
	if err != nil {
		log.Fatalf("Failed to insert endpoints: %s", err.Error())
	}
}

// CreateAndInsertArtists creates and inserts artist entries and returns it.
func (s *SeedSvc) CreateAndInsertArtists() []*models.ArtistEntry {
	dummyArtists := s.CreateArtists(s.Dummies.Artists)

	err := s.InsertArtistEntries(dummyArtists...)
	if err != nil {
		log.Fatalf("Failed to insert artist entries: %s", err.Error())
	}

	return dummyArtists
}

// CreateAndInsertArtistReferences creates and inserts artist references for given artists.
func (s *SeedSvc) CreateAndInsertArtistReferences(artists []*models.ArtistEntry) {
	artistRefs := make([]*models.ReferenceEntry, 0)

	for _, artist := range artists {
		artistRefs = append(artistRefs, s.CreateArtistReferences(artist)...)
	}

	err := s.InsertReferenceEntries(artistRefs...)
	if err != nil {
		log.Fatalf("Failed to insert artist references: %s", err.Error())
	}
}

// CreateArtistReferences creates artist references and returns it.
func (s *SeedSvc) CreateArtistReferences(artist *models.ArtistEntry) []*models.ReferenceEntry {
	artistRefs := make([]*models.ReferenceEntry, 0)

	sourceMusic := s.GetRandomSource(constants.CategoryMusic)
	artistIDRef := s.CreateReference(artist.ID, sourceMusic.ID, constants.TypeID, constants.CategoryArtist, s.CreateRandomUUID())
	artistRefs = append(artistRefs, artistIDRef)

	sourceTabs := s.GetRandomSource(constants.CategoryTabs)
	artistNameRef := s.CreateReference(artist.ID, sourceTabs.ID, constants.TypeName, constants.CategoryArtist, s.formatName(artist.Name))
	artistRefs = append(artistRefs, artistNameRef)

	return artistRefs
}

// CreateAndInsertTracks creates and inserts tracks for given artists and returns the tracks.
func (s *SeedSvc) CreateAndInsertTracks(artists []*models.ArtistEntry) []*models.TrackEntry {
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

// CreateAndInsertTrackReferences creates and inserts track references for given tracks.
func (s *SeedSvc) CreateAndInsertTrackReferences(tracks []*models.TrackEntry) {
	trackRefs := make([]*models.ReferenceEntry, 0)

	for _, track := range tracks {
		trackRefs = append(trackRefs, s.CreateTrackReferences(track)...)
	}

	err := s.InsertReferenceEntries(trackRefs...)
	if err != nil {
		log.Fatalf("Failed to insert track references: %s", err.Error())
	}
}

// CreateTrackReferences creates track references for a given track and returns it.
func (s *SeedSvc) CreateTrackReferences(track *models.TrackEntry) []*models.ReferenceEntry {
	trackRefs := make([]*models.ReferenceEntry, 0)

	sourceMusic := s.GetRandomSource(constants.CategoryMusic)
	trackIDRef := s.CreateReference(track.ID, sourceMusic.ID, constants.TypeID, constants.CategoryTrack, s.CreateRandomUUID())
	trackRefs = append(trackRefs, trackIDRef)

	sourceTabs := s.GetRandomSource(constants.CategoryTabs)
	trackNameRef := s.CreateReference(track.ID, sourceTabs.ID, constants.TypeName, constants.CategoryTrack, s.formatName(track.Title))
	trackRefs = append(trackRefs, trackNameRef)

	return trackRefs
}

// CreateAndInsertTabs creates and inserts tabs for given tacks and returns the tabs.
func (s *SeedSvc) CreateAndInsertTabs(tracks []*models.TrackEntry) []*models.TabEntry {
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

// CreateAndInsertTabReferences creates and inserts tab references for given tabs.
func (s *SeedSvc) CreateAndInsertTabReferences(tabs []*models.TabEntry) {
	tabRefs := make([]*models.ReferenceEntry, 0)

	for _, tab := range tabs {
		tabRefs = append(tabRefs, s.CreateTabReferences(tab)...)
	}

	err := s.InsertReferenceEntries(tabRefs...)
	if err != nil {
		log.Fatalf("Failed to insert tab references: %s", err.Error())
	}
}

// CreateTabReferences creates tab references for a given tab and returns it.
func (s *SeedSvc) CreateTabReferences(tab *models.TabEntry) []*models.ReferenceEntry {
	tabRefs := make([]*models.ReferenceEntry, 0)

	sourceTabs := s.GetRandomSource(constants.CategoryTabs)

	tabIDRef := s.CreateReference(tab.ID, sourceTabs.ID, constants.TypeID, constants.CategoryTab, s.CreateRandomUUID())
	tabRefs = append(tabRefs, tabIDRef)

	tabNameRef := s.CreateReference(tab.ID, sourceTabs.ID, constants.TypeName, constants.CategoryTab, s.formatName(tab.Description))
	tabRefs = append(tabRefs, tabNameRef)

	return tabRefs
}

// formatName formats the provided name.
func (s *SeedSvc) formatName(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
