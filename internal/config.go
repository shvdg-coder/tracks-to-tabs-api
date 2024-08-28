package internal

import (
	faker "github.com/brianvoe/gofakeit/v7"
	diff "github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"gopkg.in/yaml.v2"
	"os"
)

// Config represents the configuration object for the application.
type Config struct {
	Seeding *SeedingConfig `yaml:"seeding"`
}

// SeedingConfig represents the configuration with predefined seeds.
type SeedingConfig struct {
	Dummies      *DummiesConfig          `yaml:"dummies"`
	Instruments  []*diff.InstrumentEntry `yaml:"instruments"`
	Difficulties []*diff.DifficultyEntry `yaml:"difficulties"`
	Sources      []*diff.Source          `yaml:"sources"`
	Endpoints    []*diff.EndpointEntry   `yaml:"endpoints"`
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

func (a *ArtistsConfig) RandomAmount() uint {
	return uint(faker.Number(a.Min, a.Max))
}

// TracksConfig represents a struct for seeding tracks.
type TracksConfig struct {
	Min  int         `yaml:"min"`
	Max  int         `yaml:"max"`
	Tabs *TabsConfig `yaml:"tabs"`
}

func (t *TracksConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// TabsConfig represents a struct for seeing tabs.
type TabsConfig struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

func (t *TabsConfig) RandomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// NewConfig reads a file from the given path and unmarshalls its contents into a Config struct.
func NewConfig(path string) (*Config, error) {
	var config Config

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
