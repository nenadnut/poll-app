package db

import (
	"context"
	"fmt"
	"poll-app/ent"
	"poll-app/ent/question"
)

// QuestionPersistence is a concrete implementation type of the QuestionRepository
type QuestionPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *QuestionPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

// FindByID finds a question by its id
func (db *QuestionPersistence) FindByID(id int) (*ent.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	question, err := db.Client().Question.Query().
		Where(question.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching the question by id %d: %w", id, err)
	}

	return question, nil
}

// UpdateByID updates a question by its id
func (db *QuestionPersistence) UpdateByID(id int, questionUpdate *ent.QuestionUpdateOne) (*ent.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	question, err := questionUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating the question by id %d: %w", id, err)
	}

	return question, nil
}

// DeleteByID deletes a question by its id
func (db *QuestionPersistence) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	return db.Client().Question.
		DeleteOneID(id).
		Exec(ctx)
}
