package services

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"

// EndpointOps represents operations related to endpoints.
type EndpointOps interface {
	data.EndpointsData
}

// EndpointSvc is responsible for managing endpoints.
type EndpointSvc struct {
	data.EndpointsData
}

// NewEndpointSvc instantiates a new EndpointSvc.
func NewEndpointSvc(data data.EndpointsData) EndpointOps {
	return &EndpointSvc{EndpointsData: data}
}
