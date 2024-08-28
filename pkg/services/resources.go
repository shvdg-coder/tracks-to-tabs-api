package services

import (
	"fmt"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
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
		artist.Resources = r.createResourcesFromReferences("artist", artist.References...)
		r.LoadTracksResources(artist.Tracks...)
	}
}

// LoadTracksResources sets the models.Resource's for models.Track's and it's entities.
func (r *ResourceSvc) LoadTracksResources(tracks ...*models.Track) {
	for _, track := range tracks {
		references := track.References
		references = append(references, track.Artist.References...)
		track.Resources = r.createResourcesFromReferences("track", references...)
		r.LoadTabsResources(track.Tabs...)
	}
}

// LoadTabsResources sets the models.Resource's for models.Tab's.
func (r *ResourceSvc) LoadTabsResources(tabs ...*models.Tab) {
	for _, tab := range tabs {
		references := tab.References
		references = append(references, tab.Track.References...)
		references = append(references, tab.Track.Artist.References...)
		tab.Resources = r.createResourcesFromReferences("tab", references...)
	}
}

// createResourcesFromReferences creates models.Resource's for all the models.Endpoint`s that have the provided category, using the entities models.Reference's.
func (r *ResourceSvc) createResourcesFromReferences(category string, references ...*models.Reference) []*models.Resource {
	resources := make([]*models.Resource, 0)

	replacements := r.SetDefaultReplacements(r.CreateReplacements(references...))

	referencesMap := r.GroupReferencesBySource(references)
	for source, _ := range referencesMap {
		endpoints := r.FilterEndpointsByCategory(category, source.Endpoints)
		entityResources := r.createResourcesFromEndpoints(replacements, endpoints)
		resources = append(resources, entityResources...)
	}

	return resources
}

// SetDefaultReplacements sets the default replacements in a map where the key is the placeholder and the value a reference, to be used for formatting strings.
func (r *ResourceSvc) SetDefaultReplacements(replacements map[string]string) map[string]string {
	replacements["{from}"] = "0"
	replacements["{size}"] = "50"
	return replacements
}

// CreateReplacements creates createResourcesFromEndpoints map where the key is createResourcesFromEndpoints placeholder and the value a reference, to be used for formatting strings.
func (r *ResourceSvc) CreateReplacements(references ...*models.Reference) map[string]string {
	replacements := make(map[string]string)
	for _, reference := range references {
		placeholder := fmt.Sprintf("{%s:%s}", reference.Category, reference.Type)
		replacements[placeholder] = reference.Reference
	}
	return replacements
}

// GroupReferencesBySource transforms createResourcesFromEndpoints slice of models.Reference's into createResourcesFromEndpoints map where the key is the models.Source and the value createResourcesFromEndpoints slice of models.Reference's.
func (r *ResourceSvc) GroupReferencesBySource(references []*models.Reference) map[*models.Source][]*models.Reference {
	referencesMap := make(map[*models.Source][]*models.Reference)
	for _, reference := range references {
		referencesMap[reference.Source] = append(referencesMap[reference.Source], reference)
	}
	return referencesMap
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

// createResourcesFromEndpoints creates models.Resource's by formatting the URL's from the models.Endpoint using the provided replacements map.
func (r *ResourceSvc) createResourcesFromEndpoints(replacements map[string]string, endpoints []*models.Endpoint) []*models.Resource {
	resources := make([]*models.Resource, 0)
	for _, endpoint := range endpoints {
		resource := &models.Resource{Endpoint: endpoint}
		resource.FormatURL(replacements)
		resources = append(resources, resource)
	}
	return resources
}
