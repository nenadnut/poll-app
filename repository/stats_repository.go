package repository

import "poll-app/ent"

// StatsRepository is used for various aggegate operations poll voting
type StatsRepository interface {
	Client() *ent.Client
	GetPollStats(pollID int)
}
