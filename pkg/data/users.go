package data

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// UserData represents operations related to users in the database.
type UserData interface {
	InsertUser(email, password string)
	IsPasswordCorrect(username, password string) bool
}

// UserSvc is for managing users.
type UserSvc struct {
	logic.DbOperations
}

// NewUserSvc creates a new instance of the UserSvc struct.
func NewUserSvc(database logic.DbOperations) UserData {
	return &UserSvc{database}
}

// InsertUser inserts a new user into the users table.
func (d *UserSvc) InsertUser(email, plainPassword string) {
	hashedPassword, _ := logic.HashPassword(plainPassword)
	_, err := d.Exec(queries.InsertUser, email, hashedPassword)
	if err != nil {
		log.Printf("Failed inserting: %s", err.Error())
	}
}

// IsPasswordCorrect checks if the given password is correct for the user with the given email.
func (d *UserSvc) IsPasswordCorrect(email, plainPassword string) bool {
	if email == "" || plainPassword == "" {
		return false
	}
	var foundHashedPassword string
	err := d.QueryRow(queries.SelectUserPassword, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return logic.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
