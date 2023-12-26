package repository

import (
	"poll-app/ent"
)

// PollRepository is used for various persistence operations of the ent.Poll entity
type PollRepository interface {
	Client() *ent.Client
	Save(pollData *ent.PollCreate) (*ent.Poll, error)
	FindAll() ([]*ent.Poll, error)
	FindByCreator(creatorID int) ([]*ent.Poll, error)
	FindByID(id int) (*ent.Poll, error)
	UpdateByID(id int, pollUpdate *ent.PollUpdateOne) (*ent.Poll, error)
	DeleteByID(id int) error
}
