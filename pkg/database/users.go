package database

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// UserOps represents operations related to users in the database.
type UserOps interface {
	InsertUser(email, password string)
	IsPasswordCorrect(username, password string) bool
}

// UserSvc is for managing users.
type UserSvc struct {
	logic.DbOperations
}

// NewUserSvc creates a new instance of the UserSvc struct.
func NewUserSvc(database logic.DbOperations) UserOps {
	return &UserSvc{database}
}

// InsertUser inserts a new user into the users table.
func (d *UserSvc) InsertUser(email, plainPassword string) {
	hashedPassword, _ := logic.HashPassword(plainPassword)
	_, err := d.Exec(queries.InsertUserQuery, email, hashedPassword)
	if err != nil {
		log.Printf("Failed inserting user with email '%s': %s", email, err.Error())
	} else {
		log.Printf("Successfully inserted the user '%s' in the 'users' table", email)
	}
}

// IsPasswordCorrect checks if the given password is correct for the user with the given email.
func (d *UserSvc) IsPasswordCorrect(email, plainPassword string) bool {
	if email == "" || plainPassword == "" {
		return false
	}
	var foundHashedPassword string
	err := d.QueryRow(queries.SelectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return logic.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
