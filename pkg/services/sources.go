package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// SourceOps represents operations related to sources.
type SourceOps interface {
	data.SourceData
	mappers.SourceMapper
	GetSources(sourceID ...uint) ([]*models.Source, error)
}

// SourceSvc is responsible for managing sources.
type SourceSvc struct {
	data.SourceData
	mappers.SourceMapper
	EndpointOps
}

// NewSourceSvc instantiates a new SourceSvc.
func NewSourceSvc(data data.SourceData, mapper mappers.SourceMapper, endpoints EndpointOps) SourceOps {
	return &SourceSvc{SourceData: data, SourceMapper: mapper, EndpointOps: endpoints}
}

// GetSources retrieves sources with their entity references.
func (s *SourceSvc) GetSources(sourceID ...uint) ([]*models.Source, error) {
	return nil, nil
}
