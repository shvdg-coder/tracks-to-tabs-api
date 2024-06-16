package users

import "github.com/google/uuid"

// User represents a user.
type User struct {
	ID    uuid.UUID
	Email string
}

// NewUser instantiates a new User.
func NewUser(email string) *User {
	return &User{
		ID:    uuid.New(),
		Email: email,
	}
}
