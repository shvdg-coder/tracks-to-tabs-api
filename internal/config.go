package internal

import (
	faker "github.com/brianvoe/gofakeit/v7"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"gopkg.in/yaml.v2"
	"os"
)

// Config represents the configuration object for the application.
type Config struct {
	Seeding *Seeding `yaml:"seeding"`
}

// Seeding represents the configuration with predefined seeds.
type Seeding struct {
	Dummies      *Dummies           `yaml:"dummies"`
	Instruments  []*inst.Instrument `yaml:"instruments"`
	Difficulties []*diff.Difficulty `yaml:"difficulties"`
	Sources      []*src.Source      `yaml:"sources"`
	Endpoints    []*end.Endpoint    `yaml:"endpoints"`
}

// Dummies represents the configuration for generating dummies.
type Dummies struct {
	Artists *Artists `yaml:"artists"`
}

// Artists represents a struct for seeding artists.
type Artists struct {
	Min    int     `yaml:"min"`
	Max    int     `yaml:"max"`
	Tracks *Tracks `yaml:"tracks"`
}

func (a *Artists) randomAmount() uint {
	return uint(faker.Number(a.Min, a.Max))
}

// Tracks represents a struct for seeding tracks.
type Tracks struct {
	Min  int   `yaml:"min"`
	Max  int   `yaml:"max"`
	Tabs *Tabs `yaml:"tabs"`
}

func (t *Tracks) randomAmount() uint {
	return uint(faker.Number(t.Min, t.Max))
}

// Tabs represents a struct for seeing tabs.
type Tabs struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

func (t *Tabs) randomAmount() uint {
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
