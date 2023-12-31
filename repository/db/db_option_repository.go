package db

import (
	"context"
	"fmt"
	"poll-app/ent"
	"poll-app/ent/questionoption"
)

// OptionPersistence is a concrete implementation type of the OptionRepository
type OptionPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *OptionPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

// UpdateByID updates a poll by its id
func (db *OptionPersistence) UpdateByID(id int, optionUpdate *ent.QuestionOptionUpdateOne) (*ent.QuestionOption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	option, err := optionUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating the option by id %d: %w", id, err)
	}

	return option, nil
}

// DeleteByID deletes a poll by its id
func (db *OptionPersistence) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	return db.Client().QuestionOption.
		DeleteOneID(id).
		Exec(ctx)
}

// FindByQuestionID find all options that belong to some question id
func (db *OptionPersistence) FindByQuestionID(questionID int) ([]*ent.QuestionOption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	options, err := db.Client().QuestionOption.Query().
		Where(questionoption.QuestionID(questionID)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed fetching the options of the following question %d: %w", questionID, err)
	}

	return options, nil
}
