package db

import (
	"context"
	"fmt"
	"log"
	"poll-app/ent"
	"poll-app/ent/user"
	"time"
)

// UserPersistence is a concrete implementation type of the UserRepository
type UserPersistence struct {
	PersistenceClient *ent.Client
}

const dbTimeout = 3 * time.Second

// Client returns an ent database client
func (db *UserPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

// Save stores the user in a database.
func (db *UserPersistence) Save(userData *ent.UserCreate) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	user, err := userData.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating a user: %w", err)
	}

	log.Println("user has been successfully created", user)
	return user, nil
}

// FindByEmail fetches a user by email
func (db *UserPersistence) FindByEmail(email string) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	user, err := db.Client().User.Query().
		Where(user.Email(email)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching a user by email %s: %w", email, err)
	}

	return user, nil
}

// FindByID fetches a user by id
func (db *UserPersistence) FindByID(id int) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	user, err := db.Client().User.Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching a user by id %d: %w", id, err)
	}

	return user, nil
}
