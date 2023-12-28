package db

import (
	"context"
	"poll-app/ent"
	"poll-app/ent/completedquestion"
	"poll-app/ent/poll"
	"poll-app/ent/questionoption"
	"poll-app/ent/startedpoll"
	"poll-app/ent/user"
	"poll-app/internal/models"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
)

// StatsPersistence is a concrete implementation type of the StatsRepository
type StatsPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *StatsPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

func (db *StatsPersistence) GetPollStats(pollID int) (*models.PollStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	poll, err := db.Client().Poll.Query().
		Where(poll.ID(pollID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	// pollTitle := poll.Title
	questionQuery := poll.QueryQuestions()
	// all questions
	questions, err := questionQuery.Order(ent.Asc("created_at")).All(ctx)
	if err != nil {
		return nil, err
	}

	// all options
	options, err := questionQuery.QueryOptions().All(ctx)
	if err != nil {
		return nil, err
	}

	// answered questions
	answeredQuestions, err := poll.
		QueryStartedPolls().
		Where(startedpoll.Completed(true)).
		QueryCompletedQuestions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	// questionId -> option -> counter
	statsMap := make(map[int]map[int]models.OptionStats)
	for _, option := range options {
		if _, ok := statsMap[option.QuestionID]; !ok {
			statsMap[option.QuestionID] = make(map[int]models.OptionStats)
		}

		statsMap[option.QuestionID][option.ID] = models.OptionStats{
			ID:    option.ID,
			Text:  option.Text,
			Count: 0,
		}
	}

	// make a counting updates
	for _, question := range answeredQuestions {
		for _, answer := range question.Answers {
			if entry, ok := statsMap[question.QuestionID][answer]; ok {
				entry.Count++
				statsMap[question.QuestionID][answer] = entry
			}
		}
	}

	var questionStats []models.QuestionStats
	for _, question := range questions {
		var optionsStats []models.OptionStats
		for _, option := range statsMap[question.ID] {
			optionsStats = append(optionsStats, option)
		}

		questionStats = append(questionStats, models.QuestionStats{
			ID:      question.ID,
			Title:   question.Title,
			Options: optionsStats,
		})
	}

	var statsResponse models.PollStats
	statsResponse.ID = poll.ID
	statsResponse.Title = poll.Title
	statsResponse.Votes = questionStats

	return &statsResponse, nil
}

func (db *StatsPersistence) GetOptionStats(optionID int) (*models.OptionUserVoteStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var userInfos []models.OptionUserVote
	option, err := db.Client().QuestionOption.Query().
		Where(questionoption.ID(optionID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	err = db.Client().CompletedQuestion.Query().
		Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(completedquestion.FieldAnswers, optionID))
		}).
		QueryStartedPoll().
		QueryUser().
		Select(user.FieldID, user.FieldFirstName, user.FieldLastName).
		Scan(ctx, &userInfos)

	if err != nil {
		return nil, err
	}

	return &models.OptionUserVoteStats{
		ID:    optionID,
		Text:  option.Text,
		Votes: userInfos,
	}, nil
}
