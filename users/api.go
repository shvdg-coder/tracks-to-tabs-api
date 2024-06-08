package users

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"

	"log"
)

// API is for managing users.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateUsersTable creates a users table if it doesn't already exist.
func (a *API) CreateUsersTable() {
	_, err := a.Database.DB.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertUser inserts a new user into the users table.
func (a *API) InsertUser(email, plainPassword string) {
	hashedPassword, _ := logic.HashPassword(plainPassword)
	_, err := a.Database.DB.Exec(insertUserQuery, email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// IsPasswordCorrect checks if the given password is correct for the user with the given email.
func (a *API) IsPasswordCorrect(email, plainPassword string) bool {
	if email == "" || plainPassword == "" {
		return false
	}
	var foundHashedPassword string
	err := a.Database.DB.QueryRow(selectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return logic.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
