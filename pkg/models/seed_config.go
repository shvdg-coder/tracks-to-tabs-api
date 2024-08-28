package models

import (
	faker "github.com/brianvoe/gofakeit/v7"
	"gopkg.in/yaml.v2"
	"os"
)

// SeedConfig represents the configuration with seeding info.
type SeedConfig struct {
	Dummies      *DummiesConfig     `yaml:"dummies"`
	Instruments  []*InstrumentEntry `yaml:"instruments"`
	Difficulties []*DifficultyEntry `yaml:"difficulties"`
	Sources      []*SourceEntry     `yaml:"sources"`
	Endpoints    []*EndpointEntry   `yaml:"endpoints"`
}

// DummiesConfig represents the configuration for generating dummies.
type DummiesConfig struct {
	Artists *ArtistsConfig `yaml:"artists"`
}

// ArtistsConfig represents a struct for seeding artists.
type ArtistsConfig struct {
	Min    int           `yaml:"min"`
	Max    int           `yaml:"max"`
	Tracks *TracksConfig `yaml:"tracks"`
}

// RandomAmount returns a random number between the defined min and max value.
func (a *ArtistsConfig) RandomAmount() uint {
	return uint(faker.Number(a.Min, a.Max))
}

// TracksConfig represents a struct for seeding tracks.
type TracksConfig struct {
	Min  int         `yaml:"min"`
	Max  int         `yaml:"max"`
	Tabs *TabsConfig `yaml:"tabs"`
}

// RandomAmount returns a random number between the defined min and max value.
func (t *TracksConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// TabsConfig represents a struct for seeing tabs.
type TabsConfig struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// RandomAmount returns a random number between the defined min and max value.
func (t *TabsConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// NewSeedConfig reads the seed config from the given path and unmarshalls its contents into a models.SeedConfig.
func NewSeedConfig(path string) (*SeedConfig, error) {
	var config SeedConfig

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
