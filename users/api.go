package users

import (
	_ "github.com/lib/pq"
	"github.com/shvdg-dev/base-pkg/database"
	"github.com/shvdg-dev/base-pkg/utils"
	"log"
)

// API is for managing users.
type API struct {
	Database *database.Manager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *database.Manager) *API {
	return &API{Database: database}
}

// CreateUsersTable creates a users table if it doesn't already exist.
func (u *API) CreateUsersTable() {
	_, err := u.Database.DB.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertUser inserts a new user into the users table.
func (u *API) InsertUser(email, plainPassword string) {
	hashedPassword, _ := utils.HashPassword(plainPassword)
	_, err := u.Database.DB.Exec(insertUserQuery, email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// IsPasswordCorrect checks if the given password is correct for the user with the given email.
func (u *API) IsPasswordCorrect(email, plainPassword string) bool {
	if email == "" || plainPassword == "" {
		return false
	}
	var foundHashedPassword string
	err := u.Database.DB.QueryRow(selectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return utils.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
