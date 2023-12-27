package repository

import "poll-app/ent"

// StartedPollRepository is used for various persistence operations of the ent.StartedPoll entity
type StartedPollRepository interface {
	Client() *ent.Client
	Save(pollData *ent.PollCreate) (*ent.Poll, error)
	FindAll() ([]*ent.Poll, error)
	FindByCreator(creatorID int) ([]*ent.Poll, error)
	FindByID(id int) (*ent.Poll, error)
	CompletePoll(id int) (*ent.Poll, error)
	DeleteByID(id int) error
}
