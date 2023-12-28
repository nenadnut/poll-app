package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Middelware is the component that is executed on every request

func (app *application) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://*")

		if r.Method == "OPTIONS" { // preflight request
			w.Header().Set("Access-Control-Allow-Credentials", "true") // for auth
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (app *application) authRequired(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, _, err := app.auth.VerifyToken(w, r)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r, ps)
	}
}
