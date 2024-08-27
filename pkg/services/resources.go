package services

import (
	"fmt"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ResourceOps represents operations related to resources.
type ResourceOps interface {
	LoadArtistsResourcesCascading(artist ...*models.Artist)
	LoadTracksResourcesCascading(track ...*models.Track)
	LoadTabsResourcesCascading(tab ...*models.Tab)
}

// ResourceSvc is responsible for managing resources.
type ResourceSvc struct {
}

// NewResourceSvc instantiates a new ResourceSvc.
func NewResourceSvc() ResourceOps {
	return &ResourceSvc{}
}

// LoadArtistsResourcesCascading sets the models.Resource's for models.Artist's and its entities.
func (r *ResourceSvc) LoadArtistsResourcesCascading(artists ...*models.Artist) {
	for _, artist := range artists {
		artist.Resources = r.createResources(artist.References...)
		r.LoadTracksResourcesCascading(artist.Tracks...)
	}
}

// LoadTracksResourcesCascading sets the models.Resource's for models.Track's and its entities.
func (r *ResourceSvc) LoadTracksResourcesCascading(tracks ...*models.Track) {
	for _, track := range tracks {
		track.Resources = r.createResources(track.References...)
		r.LoadTabsResourcesCascading(track.Tabs...)
	}
}

// LoadTabsResourcesCascading sets the models.Resource's for models.Tab's and its entities.
func (r *ResourceSvc) LoadTabsResourcesCascading(tabs ...*models.Tab) {
	for _, tab := range tabs {
		tab.Resources = r.createResources(tab.References...)
	}
}

// createResources creates the models.Resource's from the models.Reference's.
func (r *ResourceSvc) createResources(references ...*models.Reference) []*models.Resource {
	resources := make([]*models.Resource, 0)

	referencesMap := r.GroupReferencesBySource(references)
	for source, referenceCollection := range referencesMap {
		replacements := make(map[string]string)
		for _, reference := range referenceCollection {
			placeholder := fmt.Sprintf("{%s:%s}", reference.Category, reference.Type)
			replacements[placeholder] = reference.Reference
		}
		for _, endpoint := range source.Endpoints {
			resource := &models.Resource{Endpoint: endpoint}
			resource.FormatURL(replacements)
			resources = append(resources, resource)
		}
	}

	return resources
}

// GroupReferencesBySource transforms a slice of models.Reference's into a map where the key is the models.Source and the value a slice of models.Reference's.
func (r *ResourceSvc) GroupReferencesBySource(references []*models.Reference) map[*models.Source][]*models.Reference {
	referencesMap := make(map[*models.Source][]*models.Reference, len(references))
	for _, reference := range references {
		referencesMap[reference.Source] = append(referencesMap[reference.Source], reference)
	}
	return referencesMap
}
