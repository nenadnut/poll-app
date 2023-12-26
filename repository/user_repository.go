package repository

import (
	"poll-app/ent"
)

// UserRepository is used for various persistence operations of the ent.User entity
type UserRepository interface {
	Client() *ent.Client
	Save(userData *ent.UserCreate) (*ent.User, error)
	FindByEmail(email string) (*ent.User, error)
}
