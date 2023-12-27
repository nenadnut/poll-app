package repository

import "poll-app/ent"

// CompletedQuestionRepository is used for various persistence operations of the ent.CompletedQuestion entity
type CompletedQuestionRepository interface {
	Client() *ent.Client
	Save(pollData *ent.PollCreate) (*ent.Poll, error)
	FindByStartedPollID(startedPollID int) ([]*ent.CompletedQuestion, error)
	UpdateByID(id int, completedQuestionUpdate *ent.CompletedQuestionUpdateOne) (*ent.CompletedQuestion, error)
	DeleteByID(id int) error
}
