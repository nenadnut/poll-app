package db

import (
	"context"
	"fmt"
	"log"
	"poll-app/ent"
	"poll-app/ent/poll"
)

// PollPersistence is a concrete implementation type of the PollRepository
type PollPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *PollPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

// Save stores the poll in a database.
func (db *PollPersistence) Save(pollData *ent.PollCreate) (*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	poll, err := pollData.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating a poll: %w", err)
	}

	log.Println("Poll has been successfully created: ", poll)
	return poll, nil
}

// FindAll lists all polls
func (db *PollPersistence) FindAll() ([]*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	polls, err := db.Client().Poll.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed listing the polls: %w", err)
	}

	return polls, nil
}

// FindByCreator lists all polls made by a specific creator
func (db *PollPersistence) FindByCreator(creatorID int) ([]*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	polls, err := db.Client().Poll.Query().
		Where(poll.CreatorID(creatorID)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching the polls made by user %d: %w", creatorID, err)
	}

	return polls, nil
}

// FindByID finds a poll by its id
func (db *PollPersistence) FindByID(id int) (*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	poll, err := db.Client().Poll.Query().
		Where(poll.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching the poll by id %d: %w", id, err)
	}

	return poll, nil
}

// UpdateByID updates a poll by its id
func (db *PollPersistence) UpdateByID(id int, pollUpdate *ent.PollUpdateOne) (*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	poll, err := pollUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating the poll by id %d: %w", id, err)
	}

	return poll, nil
}

// DeleteByID deletes a poll by its id
func (db *PollPersistence) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	return db.Client().Poll.
		DeleteOneID(id).
		Exec(ctx)
}
