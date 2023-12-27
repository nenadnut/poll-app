package db

import (
	"context"
	"fmt"
	"log"
	"poll-app/ent"
	"poll-app/ent/completedquestion"
)

// CompletedQuestionPersistence is a concrete implementation type of the CompletedQuestionRepository
type CompletedQuestionPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *CompletedQuestionPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

// Save stores the completed question in a database.
func (db *CompletedQuestionPersistence) Save(completedQuestionData *ent.CompletedQuestionCreate) (*ent.CompletedQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	poll, err := completedQuestionData.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed making a completed question: %w", err)
	}

	log.Println("Completed question has been successfully created: ", poll)
	return poll, nil
}

// FindByStartedPollID finds completed questions by the started poll
func (db *CompletedQuestionPersistence) FindByStartedPollID(startedPollID int) ([]*ent.CompletedQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	completedQuestions, err := db.Client().CompletedQuestion.
		Query().
		Where(completedquestion.StartedPollID(startedPollID)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching all completed questions of a started poll: %w", err)
	}

	return completedQuestions, nil
}

// UpdateByID updates a completed question by its id
func (db *CompletedQuestionPersistence) UpdateByID(id int, completedQuestionUpdate *ent.CompletedQuestionUpdateOne) (*ent.CompletedQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	completedQuestion, err := completedQuestionUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating the completed question by id %d: %w", id, err)
	}

	return completedQuestion, nil

}
