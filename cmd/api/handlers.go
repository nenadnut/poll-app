package main

import (
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

	pollCreate := app.PersistenceContext.PollPersistence.Client().Poll.
		Create().
		SetTitle(requestPayload.Title).
		SetDescription(requestPayload.Description).
		SetCreatorID(1)

	poll, err := app.PersistenceContext.PollPersistence.Save(pollCreate)
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

	polls, err := app.PersistenceContext.PollPersistence.FindAll()
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

	poll, err := app.PersistenceContext.PollPersistence.FindByID(id)
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

	pollUpdate := app.PersistenceContext.PollPersistence.
		Client().
		Poll.
		UpdateOneID(id).
		SetTitle(requestPayload.Title).
		SetDescription(requestPayload.Description)

	poll, err := app.PersistenceContext.PollPersistence.UpdateByID(id, pollUpdate)

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

	err = app.PersistenceContext.PollPersistence.DeleteByID(id)
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
	log.Printf("Add a question to poll with id %s", params.ByName("id"))

	pollID, err := strconv.Atoi(params.ByName("pollId"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// fetch the poll by id
	_, err = app.PersistenceContext.PollPersistence.FindByID(pollID)
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
	question, err := services.CreateQuestion(app.PersistenceContext, pollID, requestPayload.Questions[0])
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

	question, err := app.PersistenceContext.QuestionPersistence.FindByID(questionID)
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

	questionUpdate := app.PersistenceContext.QuestionPersistence.
		Client().
		Question.
		UpdateOneID(id).
		SetText(requestPayload.Text).
		SetNumOfAnswers(requestPayload.NumOfAnswers)

	question, err := app.PersistenceContext.QuestionPersistence.UpdateByID(id, questionUpdate)

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

	err = app.PersistenceContext.QuestionPersistence.DeleteByID(questionID)
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
func (app *application) UpdateOption(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Updating a poll by id %s", params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("Unable to extract the id parameter", err)
		utils.ErrorJSON(w, err, 400)
		return
	}

	var requestPayload struct {
		Text   string `json:"text"`
		Chosen bool   `json:"chosen"`
	}

	err = utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	optionUpdate := app.PersistenceContext.OptionPersistence.
		Client().
		QuestionOption.
		UpdateOneID(id).
		SetText(requestPayload.Text).
		SetChosen(requestPayload.Chosen)

	poll, err := app.PersistenceContext.OptionPersistence.UpdateByID(id, optionUpdate)

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

	err = app.PersistenceContext.OptionPersistence.DeleteByID(id)
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

// list polls (all vs my polls)
// edit poll
// vote
// submit
// see poll results
// poll/{id}/results
// poll/{id}/options/{}

// login
