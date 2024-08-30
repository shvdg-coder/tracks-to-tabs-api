package services

import (
	"fmt"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
)

// ResourceOps represents operations related to resources.
type ResourceOps interface {
	LoadArtistsResources(artist ...*models.Artist)
	LoadTracksResources(track ...*models.Track)
	LoadTabsResources(tab ...*models.Tab)
}

// ResourceSvc is responsible for managing resources.
type ResourceSvc struct {
}

// NewResourceSvc instantiates a new ResourceSvc.
func NewResourceSvc() ResourceOps {
	return &ResourceSvc{}
}

// LoadArtistsResources sets the models.Resource's for models.Artist's and it's entities.
func (r *ResourceSvc) LoadArtistsResources(artists ...*models.Artist) {
	for _, artist := range artists {
		artist.Resources = r.CreateResourcesFromReferences("artist", artist.References...)
	}
}

// LoadTracksResources sets the models.Resource's for models.Track's and it's entities.
func (r *ResourceSvc) LoadTracksResources(tracks ...*models.Track) {
	for _, track := range tracks {
		references := track.References
		references = append(references, track.Artist.References...)
		track.Resources = r.CreateResourcesFromReferences("track", references...)
	}
}

// LoadTabsResources sets the models.Resource's for models.Tab's.
func (r *ResourceSvc) LoadTabsResources(tabs ...*models.Tab) {
	for _, tab := range tabs {
		references := tab.References
		references = append(references, tab.Track.References...)
		references = append(references, tab.Track.Artist.References...)
		tab.Resources = r.CreateResourcesFromReferences("tab", references...)
	}
}

// CreateResourcesFromReferences creates models.Resource's for all the models.Endpoint`s that have the provided category, using the entities models.Reference's.
func (r *ResourceSvc) CreateResourcesFromReferences(category string, references ...*models.Reference) []*models.Resource {
	resources := make([]*models.Resource, 0)

	replacements := r.SetDefaultReplacements(r.CreateReplacements(references...))

	sourceMap := r.GroupSources(references)
	for _, source := range sourceMap {
		endpoints := r.FilterEndpointsByCategory(category, source.Endpoints)
		entityResources := r.CreateResourcesFromEndpoints(replacements, endpoints)
		resources = append(resources, entityResources...)
	}

	return resources
}

// SetDefaultReplacements sets the default replacements in the provided map, where the key is the placeholder and the value a reference, to be used for formatting strings.
func (r *ResourceSvc) SetDefaultReplacements(replacements map[string]string) map[string]string {
	replacements["{from}"] = "0"
	replacements["{size}"] = "50"
	return replacements
}

// CreateReplacements creates CreateResourcesFromEndpoints map where the key is CreateResourcesFromEndpoints placeholder and the value a reference, to be used for formatting strings.
func (r *ResourceSvc) CreateReplacements(references ...*models.Reference) map[string]string {
	replacements := make(map[string]string)
	for _, reference := range references {
		placeholder := fmt.Sprintf("{%s:%s}", reference.Category, reference.Type)
		replacements[placeholder] = reference.Reference
	}
	return replacements
}

// GroupSources collect models.Source's from a slice of models.Reference's, and put them in a map where the key is the ID and the value is a models.Source.
func (r *ResourceSvc) GroupSources(references []*models.Reference) map[uint]*models.Source {
	sourceMap := make(map[uint]*models.Source)
	for _, reference := range references {
		sourceMap[reference.SourceID] = reference.Source
	}
	return sourceMap
}

// FilterEndpointsByCategory plucks the endpoints of which the category corresponds with the provided category.
func (r *ResourceSvc) FilterEndpointsByCategory(category string, endpoints []*models.Endpoint) []*models.Endpoint {
	filteredEndpoints := make([]*models.Endpoint, 0)
	for _, endpoint := range endpoints {
		if endpoint.Category == category {
			filteredEndpoints = append(filteredEndpoints, endpoint)
		}
	}
	return filteredEndpoints
}

// CreateResourcesFromEndpoints creates models.Resource's by formatting the URL's from the models.Endpoint using the provided replacements map.
func (r *ResourceSvc) CreateResourcesFromEndpoints(replacements map[string]string, endpoints []*models.Endpoint) []*models.Resource {
	resources := make([]*models.Resource, 0)
	for _, endpoint := range endpoints {
		resource := &models.Resource{Endpoint: endpoint}
		resource.FormatURL(replacements)
		resources = append(resources, resource)
	}
	return resources
}
