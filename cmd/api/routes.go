package main

import "github.com/julienschmidt/httprouter"

func (app *application) router() *httprouter.Router {

	router := httprouter.New()

	// polls
	router.POST("/polls", app.CreatePoll)
	router.GET("/polls", app.ListPolls)
	router.GET("/polls/:id", app.GetPoll)
	router.PUT("/polls/:id", app.UpdatePoll)
	router.DELETE("/polls/:id", app.DeletePoll)

	// questions
	// design flaw in httprouter - panic: ':pollId' in new path '/polls/:pollId/questions/:questionId' conflicts with
	// existing wildcard ':id' in existing prefix '/polls/:id'; it doesn't support REST-based paths
	router.POST("/polls/:pollId/questions", app.AddQuestions)
	router.GET("/questions/:id", app.GetQuestion)
	router.PUT("/questions/:id", app.UpdateQuestion)
	router.DELETE("/questions/:id", app.DeleteQuestion)

	// options
	router.PUT("/options/:id", app.UpdateOption)
	router.DELETE("/options/:id", app.DeleteOption)

	return router
}
