package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// UserSchema represents schema operations related to 'users' in the database.
type UserSchema interface {
	CreateUsersTable()
	DropUsersTable()
}

// UserSvc is for managing 'users' tables in the database.
type UserSvc struct {
	logic.DbOps
}

// NewUserSvc creates a new instance of the UserSvc struct.
func NewUserSvc(database logic.DbOps) UserSchema {
	return &UserSvc{database}
}

// CreateUsersTable creates a users table if it doesn't already exist.
func (s *UserSvc) CreateUsersTable() {
	_, err := s.DB().Exec(queries.CreateUsersTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropUsersTable drops the users table if it exists.
func (s *UserSvc) DropUsersTable() {
	_, err := s.DB().Exec(queries.DropUsersTable)
	if err != nil {
		log.Fatal(err)
	}
}
