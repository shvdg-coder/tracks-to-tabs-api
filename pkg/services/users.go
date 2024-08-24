package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// UserOps represents the operations related to users.
type UserOps interface {
	schemas.UserSchema
	data.UserData
}

// UserSvc is responsible for managing users.
type UserSvc struct {
	schemas.UserSchema
	data.UserData
}

// NewUserSvc creates a new instance of UserSvc.
func NewUserSvc(schema schemas.UserSchema, data data.UserData) UserOps {
	return &UserSvc{
		UserSchema: schema,
		UserData:   data,
	}
}
