package db

import (
	"context"
	"fmt"
	"log"
	"poll-app/ent"
	"poll-app/ent/startedpoll"
)

// StartedPollPersistence is a concrete implementation type of the StartedPollRepository
type StartedPollPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *StartedPollPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

// Save stores the started poll in a database.
func (db *StartedPollPersistence) Save(startedPollCreate *ent.StartedPollCreate) (*ent.StartedPoll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	startedPoll, err := startedPollCreate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed starting a poll: %w", err)
	}

	log.Println("Poll has been successfully created: ", startedPoll)
	return startedPoll, nil
}

// FindAll lists all polls
func (db *StartedPollPersistence) FindAll() ([]*ent.StartedPoll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	startedPolls, err := db.Client().StartedPoll.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed listing the started polls: %w", err)
	}

	return startedPolls, nil
}

// FindByCreator lists all polls made by a specific creator
func (db *StartedPollPersistence) FindByCreator(userID int) ([]*ent.StartedPoll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	startedPolls, err := db.Client().StartedPoll.Query().
		Where(startedpoll.UserID(userID)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching the started polls made by user %d: %w", userID, err)
	}

	return startedPolls, nil
}

// FindByID finds a poll by its id
func (db *StartedPollPersistence) FindByID(id int) (*ent.StartedPoll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	startedPoll, err := db.Client().StartedPoll.Query().
		Where(startedpoll.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching the started poll by id %d: %w", id, err)
	}

	return startedPoll, nil
}

// CompletePoll completes the poll with the given id
func (db *StartedPollPersistence) CompletePoll(id int) (*ent.StartedPoll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	completedPoll, err := db.Client().StartedPoll.
		UpdateOneID(id).
		SetCompleted(true).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed completing the poll by id %d: %w", id, err)
	}

	return completedPoll, nil
}

// DeleteByID deletes a started poll by its id
func (db *StartedPollPersistence) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	return db.Client().StartedPoll.
		DeleteOneID(id).
		Exec(ctx)
}
