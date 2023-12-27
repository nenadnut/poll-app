package repository

import "poll-app/ent"

// QuestionRepository is used for various persistence operations of the ent.Question entity
type QuestionRepository interface {
	Client() *ent.Client
	FindByID(id int) (*ent.Question, error)
	FindRequiredByPollID(pollID int) ([]*ent.Question, error)
	UpdateByID(id int, questionUpdate *ent.QuestionUpdateOne) (*ent.Question, error)
	DeleteByID(id int) error
}
