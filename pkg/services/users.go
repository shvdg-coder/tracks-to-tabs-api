package services

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"

// Operations represents the operations related to users.
type Operations interface {
	database.UserOps
}

// Service is responsible for managing users.
type Service struct {
	database.UserOps
}

// NewService creates a new instance of Service
func NewService(data database.UserOps) Operations {
	return &Service{UserOps: data}
}
