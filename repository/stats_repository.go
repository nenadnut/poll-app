package repository

import (
	"poll-app/ent"
	"poll-app/internal/models"
)

// StatsRepository is used for various aggegate operations poll voting
type StatsRepository interface {
	Client() *ent.Client
	GetPollStats(pollID int) (*models.PollStats, error)
	GetOptionStats(optionID int) (*models.OptionUserVoteStats, error)
}
