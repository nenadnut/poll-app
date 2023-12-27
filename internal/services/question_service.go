package services

import (
	"context"
	"poll-app/ent"
	"poll-app/internal/models"
	"poll-app/repository"
	"time"
)

// CreateQuestion creates and persists a batch of questions
func CreateQuestion(persistence *repository.PersistenceContext, pollID int, questionRequest models.QuestionRequest) (*ent.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	question, err := persistence.PollPersistence.PersistenceClient.Question.
		Create().
		SetPollID(pollID).
		SetText(questionRequest.Text).
		SetNumOfAnswers(questionRequest.NumOfAnswers).
		SetRequired(questionRequest.Required).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	options, err := CreateQuestionOptions(ctx, persistence, question.ID, questionRequest.Options)
	if err != nil {
		return nil, err
	}

	question.Edges.Options = append(question.Edges.Options, options...)
	return question, nil
}

func createQuestionBuilder(persistence *repository.PersistenceContext, pollID int, questionReq models.QuestionRequest) *ent.QuestionCreate {
	question := persistence.PollPersistence.PersistenceClient.Question.
		Create().
		SetPollID(pollID).
		SetText(questionReq.Text).
		SetNumOfAnswers(questionReq.NumOfAnswers)

	return question
}

// CreateQuestionOptions creates question options of a particular question
func CreateQuestionOptions(context context.Context, persistence *repository.PersistenceContext, questionID int, optionReqs []models.OptionRequest) ([]*ent.QuestionOption, error) {
	var optionBuilders []*ent.QuestionOptionCreate

	for _, optionReq := range optionReqs {
		optionBuilders = append(optionBuilders, persistence.UserPersistence.PersistenceClient.QuestionOption.Create().
			SetText(optionReq.Text).
			SetQuestionID(questionID))
	}

	options, err := persistence.PollPersistence.PersistenceClient.QuestionOption.CreateBulk(optionBuilders...).Save(context)
	if err != nil {
		return nil, err
	}

	return options, nil
}

// CreateQuestionOptionsWithoutContext creates question options of a particular question
func CreateQuestionOptionsWithoutContext(persistence *repository.PersistenceContext, questionID int, optionReqs []models.OptionRequest) ([]*ent.QuestionOption, error) {
	context, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var optionBuilders []*ent.QuestionOptionCreate

	for _, optionReq := range optionReqs {
		optionBuilders = append(optionBuilders, persistence.UserPersistence.PersistenceClient.QuestionOption.Create().
			SetText(optionReq.Text).
			SetQuestionID(questionID))
	}

	options, err := persistence.PollPersistence.PersistenceClient.QuestionOption.CreateBulk(optionBuilders...).Save(context)
	if err != nil {
		return nil, err
	}

	return options, nil
}
