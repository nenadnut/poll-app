package db

import (
	"context"
	"log"
	"poll-app/ent"
	"poll-app/ent/completedquestion"
	"poll-app/ent/questionoption"
	"poll-app/ent/startedpoll"
)

// StatsPersistence is a concrete implementation type of the StatsRepository
type StatsPersistence struct {
	PersistenceClient *ent.Client
}

// Client returns an ent database client
func (db *StatsPersistence) Client() *ent.Client {
	return db.PersistenceClient
}

func (db *StatsPersistence) GetPollStats(pollID int) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	completedPolls, err := db.Client().StartedPoll.Query().
		Where(startedpoll.PollID(pollID)).
		Where(startedpoll.Completed(true)).
		All(ctx)

	if err != nil {
		log.Println(err)
	}

	// extracting the poll ids
	var completedPollIDs []int
	for _, completedPoll := range completedPolls {
		completedPollIDs = append(completedPollIDs, completedPoll.ID)
	}

	answeredQuestions, err := db.Client().CompletedQuestion.Query().
		Where(completedquestion.StartedPollIDIn(completedPollIDs...)).
		All(ctx)

	// extract question ids
	var questionIds []int
	// questionTitles := make(map[int]string)

	for _, question := range answeredQuestions {
		questionIds = append(questionIds, question.QuestionID)
		// if _, ok := questionTitles[question.QuestionID]; !ok {
		// 	questionTitles[question.QuestionID] = question.QueryQuestion().Where()

		// }
	}

	log.Println(questionIds)

	// extract all possible question options
	options, err := db.Client().QuestionOption.
		Query().
		Where(questionoption.QuestionIDIn(questionIds...)).
		All(ctx)

	if err != nil {
		log.Println(err)
	}

	// questionId -> option -> counter
	statsMap := make(map[int]map[int]int)
	for _, option := range options {
		statsMap[option.QuestionID] = make(map[int]int)
		statsMap[option.QuestionID][option.ID] = 0
	}

	log.Println(statsMap)

	// make a counting updates
	for _, question := range answeredQuestions {
		for _, answer := range question.Answers {
			statsMap[question.QuestionID][answer]++
		}
	}

	log.Println(statsMap)
}
