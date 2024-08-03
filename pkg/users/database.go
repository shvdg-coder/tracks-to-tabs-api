package users

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to users in the database.
type DataOperations interface {
	InsertUser(email, password string)
	IsPasswordCorrect(username, password string) bool
}

// DataService is for managing users.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{database}
}

// InsertUser inserts a new user into the users table.
func (d *DataService) InsertUser(email, plainPassword string) {
	hashedPassword, _ := logic.HashPassword(plainPassword)
	_, err := d.Exec(insertUserQuery, email, hashedPassword)
	if err != nil {
		log.Printf("Failed inserting user with email '%s': %s", email, err.Error())
	} else {
		log.Printf("Successfully inserted the user '%s' in the 'users' table", email)
	}
}

// IsPasswordCorrect checks if the given password is correct for the user with the given email.
func (d *DataService) IsPasswordCorrect(email, plainPassword string) bool {
	if email == "" || plainPassword == "" {
		return false
	}
	var foundHashedPassword string
	err := d.QueryRow(selectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return logic.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
