package main

import (
	"fmt"
	"log"
	"net/http"
	"poll-app/internal/models"
	"poll-app/internal/services"
	"poll-app/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ////// Polls ////////

func (app *application) CreatePoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("Creating a poll...")
	var requestPayload struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	pollCreate := app.persistenceContext.PollPersistence.Client().Poll.
		Create().
		SetTitle(requestPayload.Title).
		SetDescription(requestPayload.Description).
		SetCreatorID(1)

	poll, err := app.persistenceContext.PollPersistence.Save(pollCreate)
	if err != nil {
		log.Printf("failed creating a poll: %s", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJSON(w, 201, poll)
	if err != nil {
		log.Printf("poll creation failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

// ListPolls will list all polls made by a user (all polls in case the user has an ADMIN role)
func (app *application) ListPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Listing all polls...")

	polls, err := app.persistenceContext.PollPersistence.FindAll()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, 200, polls)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

// GetPoll will fetch a specific poll or return 404
func (app *application) GetPoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Get poll by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	poll, err := app.persistenceContext.PollPersistence.FindByID(id)
	if err != nil {
		log.Printf("Could not fetch the required poll[id = %d] - %s", id, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 200, poll)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

func (app *application) UpdatePoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Updating a poll by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	var requestPayload struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err = utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	pollUpdate := app.persistenceContext.PollPersistence.
		Client().
		Poll.
		UpdateOneID(id).
		SetTitle(requestPayload.Title).
		SetDescription(requestPayload.Description)

	poll, err := app.persistenceContext.PollPersistence.UpdateByID(id, pollUpdate)

	if err != nil {
		log.Printf("Could not update the poll[id = %d] - %s", id, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 200, poll)
	if err != nil {
		log.Printf("poll creation failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

// DeletePoll will delete a specific poll or return 404
func (app *application) DeletePoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Delete poll by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	err = app.persistenceContext.PollPersistence.DeleteByID(id)
	if err != nil {
		log.Printf("Could not delete the poll[id = %d] - %s", id, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 204, nil)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

// //// Questions //////
func (app *application) AddQuestions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Add a question to poll with id %s", params.ByName("pollId"))

	pollID, err := strconv.Atoi(params.ByName("pollId"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// fetch the poll by id
	_, err = app.persistenceContext.PollPersistence.FindByID(pollID)
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	var requestPayload models.QuestionsBulkRequest
	err = utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("Unable to deserialize the request data", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Println("Received a bulk of questions to create", requestPayload)
	question, err := services.CreateQuestion(app.persistenceContext, pollID, requestPayload.Questions[0])
	if err != nil {
		log.Println("Unable to deserialize the request data", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJSON(w, 201, question)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

// GetQuestion gets a question by its ID.
func (app *application) GetQuestion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Get a question with ID  %s", params.ByName("id"))

	questionID, err := strconv.Atoi(params.ByName("questionId"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	question, err := app.persistenceContext.QuestionPersistence.FindByID(questionID)
	if err != nil {
		log.Printf("Could not fetch the required question[id = %d] - %s", questionID, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 200, question)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

func (app *application) UpdateQuestion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Updating a question by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	var requestPayload struct {
		Text         string `json:"text"`
		NumOfAnswers int    `json:"num_of_answers"`
	}

	err = utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	questionUpdate := app.persistenceContext.QuestionPersistence.
		Client().
		Question.
		UpdateOneID(id).
		SetText(requestPayload.Text).
		SetNumOfAnswers(requestPayload.NumOfAnswers)

	question, err := app.persistenceContext.QuestionPersistence.UpdateByID(id, questionUpdate)

	if err != nil {
		log.Printf("Could not update the question[id = %d] - %s", id, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 200, question)
	if err != nil {
		log.Printf("question creation failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

// DeleteQuestion deletes a question by its ID.
func (app *application) DeleteQuestion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Delete a question with ID  %s", params.ByName("id"))

	questionID, err := strconv.Atoi(params.ByName("questionId"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.persistenceContext.QuestionPersistence.DeleteByID(questionID)
	if err != nil {
		log.Printf("Could not delete the question[id = %d] - %s", questionID, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 204, nil)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

// ///// Options ///////
func (app *application) AddOptions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Add options to question with id %s", params.ByName("questionId"))

	questionID, err := strconv.Atoi(params.ByName("questionId"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// fetch the question by id
	_, err = app.persistenceContext.QuestionPersistence.FindByID(questionID)
	if err != nil {
		log.Printf("question with id %d not found", questionID)
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	var requestPayload models.OptionsBulkRequest
	err = utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("Unable to deserialize the request data", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Println("Received a bulk of options to create", requestPayload)
	options, err := services.CreateQuestionOptionsWithoutContext(app.persistenceContext, questionID, requestPayload.Options)
	if err != nil {
		log.Println("Some error occurred", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, 201, options)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

func (app *application) UpdateOption(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Updating a poll by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	var requestPayload struct {
		Text string `json:"text"`
	}

	err = utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	optionUpdate := app.persistenceContext.OptionPersistence.
		Client().
		QuestionOption.
		UpdateOneID(id).
		SetText(requestPayload.Text)

	poll, err := app.persistenceContext.OptionPersistence.UpdateByID(id, optionUpdate)

	if err != nil {
		log.Printf("could not update an option[id = %d] - %s", id, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 200, poll)
	if err != nil {
		log.Printf("option update failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

// DeleteOption will delete a specific option or return 404
func (app *application) DeleteOption(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Delete an option by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	err = app.persistenceContext.OptionPersistence.DeleteByID(id)
	if err != nil {
		log.Printf("Could not delete the option[id = %d] - %s", id, err)
		utils.ErrorJSON(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, 204, nil)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

// // Start poll ////
func (app *application) StartPoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("Starting a poll...")

	var requestPayload struct {
		PollID int `json:"poll_id"`
		UserID int `json:"user_id"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// TODO change - extract the userId from the session
	startedPollCreate := app.persistenceContext.StartedPollPersistence.Client().StartedPoll.
		Create().
		SetPollID(requestPayload.PollID).
		SetUserID(requestPayload.UserID)

	startedPoll, err := app.persistenceContext.StartedPollPersistence.Save(startedPollCreate)
	if err != nil {
		log.Printf("failed starting a poll: %s", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJSON(w, 201, startedPoll)
	if err != nil {
		log.Printf("poll creation failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

func (app *application) CompletePoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("Completing a poll...", params.ByName("pollId"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	startedPoll, err := app.persistenceContext.StartedPollPersistence.FindByID(id)
	if err != nil {
		log.Printf("failed starting a poll: %s", err)
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	// Finds all required questions of that poll
	requiredQuestions, err := app.persistenceContext.QuestionPersistence.FindRequiredByPollID(startedPoll.PollID)
	if len(requiredQuestions) > 0 {
		completedQuestions, err := app.persistenceContext.CompletedQuestionPersistence.FindByStartedPollID(startedPoll.ID)
		if err != nil {
			log.Printf("failed finding the completed questions: %s", err)
			utils.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		// compare required questions and completed questions
		answeredQuestionsMap := make(map[int]int)
		for _, answeredQuestion := range completedQuestions {
			answeredQuestionsMap[answeredQuestion.QuestionID] = answeredQuestion.ID
		}

		for _, question := range requiredQuestions {
			if _, ok := answeredQuestionsMap[question.ID]; !ok {
				err := fmt.Errorf("failed completing the poll; some mandatory questions are not answered")
				log.Printf(err.Error())
				utils.ErrorJSON(w, err, http.StatusForbidden)
				return
			}
		}
	}

	completedPoll, err := app.persistenceContext.StartedPollPersistence.CompletePoll(id)
	if err != nil {
		log.Printf("failed completing a poll: %s", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJSON(w, 200, completedPoll)
	if err != nil {
		log.Printf("poll update failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

func (app *application) AnswerQuestion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("Answering a question...")

	var requestPayload struct {
		PollID     int   `json:"started_poll_id"`
		QuestionID int   `json:"question_id"`
		Answers    []int `json:"answers"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	question, err := app.persistenceContext.QuestionPersistence.FindByID(requestPayload.QuestionID)
	if err != nil {
		log.Printf("failed finding a question: %s", err)
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	numOfAnswers := len(requestPayload.Answers)
	if numOfAnswers == 0 || numOfAnswers != question.NumOfAnswers {
		err = fmt.Errorf("failed when answering a question, this question requires exactly %d answer(s)", question.NumOfAnswers)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// check if the options actually belong to the question
	optionIds := requestPayload.Answers

	if len(optionIds) != 0 {
		questionOptions, err := app.persistenceContext.OptionPersistence.FindByQuestionID(question.ID)
		if err != nil {
			err = fmt.Errorf("failed fetching options")
			utils.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		questionOptionIDMap := make(map[int]int)
		for _, option := range questionOptions {
			questionOptionIDMap[option.ID] = option.ID
		}

		for _, optionID := range optionIds {
			if _, ok := questionOptionIDMap[optionID]; !ok {
				err := fmt.Errorf("answer contains an invalid option: %d", optionID)
				utils.ErrorJSON(w, err, http.StatusBadRequest)
				return
			}
		}
	}

	completedQuestionBuilder := app.persistenceContext.CompletedQuestionPersistence.Client().CompletedQuestion.
		Create().
		SetStartedPollID(requestPayload.PollID).
		SetQuestionID(requestPayload.QuestionID).
		SetAnswers(requestPayload.Answers)

	completedQuestion, err := app.persistenceContext.CompletedQuestionPersistence.Save(completedQuestionBuilder)
	if err != nil {
		log.Printf("failed asnwering a question: %s", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJSON(w, 201, completedQuestion)
	if err != nil {
		log.Printf("answering a question failed due to %s", err)
		utils.ErrorJSON(w, err)
		return
	}
}

func (app *application) GetPollStats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pollID := 2

	app.persistenceContext.StatsPersistence.GetPollStats(pollID)
}

/// Question stats ////

/// Per user stats who clicked on which option

//// ----

///// Question answers ////

//// See poll stats /////

//// Get result of poll /////

//// Check poll answer stats ////

// login
