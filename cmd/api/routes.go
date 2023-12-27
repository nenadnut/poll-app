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
	router.POST("/options/:questionId", app.AddOptions)
	router.PUT("/options/:id", app.UpdateOption)
	router.DELETE("/options/:id", app.DeleteOption)

	// started polls
	router.POST("/started-polls/", app.StartPoll)
	router.PUT("/started-polls/:id/complete", app.CompletePoll)

	// answered questions
	router.POST("/answered-questions/", app.AnswerQuestion)

	// stats
	router.GET("/completed-polls/:polld/stats", app.GetPollStats)

	return router
}
