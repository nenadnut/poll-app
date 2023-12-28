package main

import "github.com/julienschmidt/httprouter"

func (app *application) router() *httprouter.Router {

	router := httprouter.New()

	// polls
	router.POST("/polls", app.authRequired(app.CreatePoll))
	router.GET("/polls", app.authRequired(app.ListPolls))
	router.GET("/polls/:id", app.authRequired(app.GetPoll))
	router.PUT("/polls/:id", app.authRequired(app.UpdatePoll))
	router.DELETE("/polls/:id", app.authRequired(app.DeletePoll))

	// questions
	// design flaw in httprouter - panic: ':pollId' in new path '/polls/:pollId/questions/:questionId' conflicts with
	// existing wildcard ':id' in existing prefix '/polls/:id'; it doesn't support REST-based paths
	router.POST("/polls/:pollId/questions", app.authRequired(app.AddQuestions))
	router.GET("/questions/:id", app.authRequired(app.GetQuestion))
	router.PUT("/questions/:id", app.authRequired(app.UpdateQuestion))
	router.DELETE("/questions/:id", app.authRequired(app.DeleteQuestion))

	// options
	router.POST("/options/:questionId", app.authRequired(app.AddOptions))
	router.PUT("/options/:id", app.authRequired(app.UpdateOption))
	router.DELETE("/options/:id", app.authRequired(app.DeleteOption))

	// started polls
	router.POST("/started-polls/", app.authRequired(app.StartPoll))
	router.PUT("/started-polls/:id/complete", app.authRequired(app.CompletePoll))

	// answered questions
	router.POST("/answered-questions/", app.authRequired(app.AnswerQuestion))

	// stats
	router.GET("/completed-polls/:pollId/stats", app.authRequired(app.GetPollStats))
	// option voting stats
	router.GET("/completed-options/:optionId/stats", app.authRequired(app.GetOptionStats))

	router.POST("/login", app.authenticate)
	router.POST("/logout", app.authRequired(app.logout))
	router.POST("/refresh-token", app.authRequired(app.refreshToken))

	return router
}
