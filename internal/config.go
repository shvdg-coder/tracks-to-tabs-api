package internal

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Config represents the configuration object for the application.
type Config struct {
	Seeds Seeds `yaml:"seeds"`
}

// Seeds represents the configuration for the seed data.
type Seeds struct {
	Artists   Artists    `yaml:"artists"`
	Tracks    Tracks     `yaml:"tracks"`
	Tabs      Tabs       `yaml:"tabs"`
	Sources   []Source   `yaml:"sources"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

// Artists represents a struct for seeding artists.
type Artists struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// Tracks represents a struct for seeding tracks.
type Tracks struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// Tabs represents a struct for seeing tabs.
type Tabs struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// Source represents a struct for seeding a source.
type Source struct {
	Name string `yaml:"name"`
}

// Endpoint represents a struct for seeding an endpoint.
type Endpoint struct {
	Source   string `yaml:"source"`
	Category string `yaml:"category"`
	URL      string `yaml:"URL"`
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
