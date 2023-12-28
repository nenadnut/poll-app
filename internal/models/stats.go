package models

// OptionStats holds the stats data of the votes for the particular option
type OptionStats struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Count int    `json:"count"`
}

// QuestionStats holds the stats data of the question votes
type QuestionStats struct {
	ID      int           `json:"id"`
	Title   string        `json:"text"`
	Options []OptionStats `json:"options"`
}

// PollStats holds the stats data of the poll question votes
type PollStats struct {
	ID    int             `json:"id"`
	Title string          `json:"title"`
	Votes []QuestionStats `json:"votes"`
}

type OptionUserVoteStats struct {
	ID    int              `json:"id"`
	Text  string           `json:"text"`
	Votes []OptionUserVote `json:"votes"`
}

type OptionUserVote struct {
	ID       int    `json:"id"`
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
}
