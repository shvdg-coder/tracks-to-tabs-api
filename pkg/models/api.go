package models

import (
	faker "github.com/brianvoe/gofakeit/v7"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"gopkg.in/yaml.v3"
	"os"
)

// APIConfig holds the overall configuration for the application
type APIConfig struct {
	Commands *CommandsConfig `yaml:"commands"`
	Database *DatabaseConfig `yaml:"database"`
	Seeding  *SeedingConfig  `yaml:"seeding"`
}

// CommandsConfig holds the configuration related to the commands
type CommandsConfig struct {
	CreateEnabled bool `yaml:"create_enabled"`
	DropEnabled   bool `yaml:"drop_enabled"`
	SeedEnabled   bool `yaml:"seed_enabled"`
}

// DatabaseConfig holds the configuration related to the database
type DatabaseConfig struct {
	URL string           `yaml:"url"`
	SSH *logic.SSHConfig `yaml:"ssh"`
}

// SeedingConfig holds the configuration related to seeding the application
type SeedingConfig struct {
	Dummies      DummyConfig        `yaml:"dummies"`
	Instruments  []*InstrumentEntry `yaml:"instruments"`
	Difficulties []*DifficultyEntry `yaml:"difficulties"`
	Sources      []*SourceEntry     `yaml:"sources"`
	Endpoints    []*EndpointEntry   `yaml:"endpoints"`
}

// DummyConfig contains the configuration for creating dummy data
type DummyConfig struct {
	Artists *ArtistConfig `yaml:"artists"`
}

// ArtistConfig covers configuration related to artist data
type ArtistConfig struct {
	Min    int          `yaml:"min"`
	Max    int          `yaml:"max"`
	Tracks *TrackConfig `yaml:"tracks"`
}

// RandomAmount returns a random number between the defined min and max value.
func (a *ArtistConfig) RandomAmount() uint {
	return uint(faker.Number(a.Min, a.Max))
}

// TrackConfig represents configuration related to audio tracks
type TrackConfig struct {
	Min  int        `yaml:"min"`
	Max  int        `yaml:"max"`
	Tabs *TabConfig `yaml:"tabs"`
}

// RandomAmount returns a random number between the defined min and max value.
func (t *TrackConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// TabConfig defines configuration relevant to tabs
type TabConfig struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
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
