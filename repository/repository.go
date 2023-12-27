package repository

import (
	"poll-app/ent"
	"poll-app/repository/db"
)

// PersistenceContext is a holder for all persistence objects
type PersistenceContext struct {
	UserPersistence              db.UserPersistence
	PollPersistence              db.PollPersistence
	QuestionPersistence          db.QuestionPersistence
	OptionPersistence            db.OptionPersistence
	StartedPollPersistence       db.StartedPollPersistence
	CompletedQuestionPersistence db.CompletedQuestionPersistence
	StatsPersistence             db.StatsPersistence
}

// New constructs a new PersistenceContext
func New(client *ent.Client) *PersistenceContext {
	return &PersistenceContext{
		UserPersistence:              db.UserPersistence{PersistenceClient: client},
		PollPersistence:              db.PollPersistence{PersistenceClient: client},
		QuestionPersistence:          db.QuestionPersistence{PersistenceClient: client},
		OptionPersistence:            db.OptionPersistence{PersistenceClient: client},
		StartedPollPersistence:       db.StartedPollPersistence{PersistenceClient: client},
		CompletedQuestionPersistence: db.CompletedQuestionPersistence{PersistenceClient: client},
		StatsPersistence:             db.StatsPersistence{PersistenceClient: client},
	}
}
