package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
)

// DifficultyOps represents operations related to difficulties.
type DifficultyOps interface {
	data.DifficultyData
	mappers.DifficultyMapper
}

// DifficultySvc is responsible for managing difficulties.
type DifficultySvc struct {
	data.DifficultyData
	mappers.DifficultyMapper
}

// NewDifficultySvc instantiates a new instance of DifficultySvc.
func NewDifficultySvc(data data.DifficultyData, mapper mappers.DifficultyMapper) DifficultyOps {
	return &DifficultySvc{DifficultyData: data, DifficultyMapper: mapper}
}
