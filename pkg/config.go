package pkg

import (
	faker "github.com/brianvoe/gofakeit/v7"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"gopkg.in/yaml.v3"
	"os"
)

// APIConfig holds the overall configuration for the application
type APIConfig struct {
	Database *DatabaseConfig `json:"database"`
	Seeding  *SeedingConfig  `json:"seeding"`
}

// DatabaseConfig holds the configuration related to the database
type DatabaseConfig struct {
	URL string           `json:"url"`
	SSH *logic.SSHConfig `json:"ssh"`
}

// SeedingConfig holds the configuration related to seeding the application
type SeedingConfig struct {
	Dummies      DummyConfig               `json:"dummies"`
	Instruments  []*models.InstrumentEntry `json:"instruments"`
	Difficulties []*models.DifficultyEntry `json:"difficulties"`
	Sources      []*models.SourceEntry     `json:"sources"`
	Endpoints    []*models.EndpointEntry   `json:"endpoints"`
}

// DummyConfig contains the configuration for creating dummy data
type DummyConfig struct {
	Artists *ArtistConfig `json:"artists"`
}

// ArtistConfig covers configuration related to artist data
type ArtistConfig struct {
	Min    int          `json:"min"`
	Max    int          `json:"max"`
	Tracks *TrackConfig `json:"tracks"`
}

// RandomAmount returns a random number between the defined min and max value.
func (a *ArtistConfig) RandomAmount() uint {
	return uint(faker.Number(a.Min, a.Max))
}

// TrackConfig represents configuration related to audio tracks
type TrackConfig struct {
	Min  int        `json:"min"`
	Max  int        `json:"max"`
	Tabs *TabConfig `json:"tabs"`
}

// RandomAmount returns a random number between the defined min and max value.
func (t *TrackConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// TabConfig defines configuration relevant to tabs
type TabConfig struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// RandomAmount returns a random number between the defined min and max value.
func (t *TabConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// NewAPIConfig reads the API config from the path and unmarshalls its contents into a APIConfig.
func NewAPIConfig(path string) (*APIConfig, error) {
	var config APIConfig

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
