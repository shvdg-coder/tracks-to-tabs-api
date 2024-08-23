package services

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"

// UserOps represents the operations related to users.
type UserOps interface {
	data.UserData
}

// UserSvc is responsible for managing users.
type UserSvc struct {
	data.UserData
}

// NewUserSvc creates a new instance of UserSvc.
func NewUserSvc(data data.UserData) UserOps {
	return &UserSvc{UserData: data}
}
