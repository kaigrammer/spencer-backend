package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// API health status endpoint
	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)

	// Expenses endpoint
	router.HandlerFunc(http.MethodPost, "/api/v1/expenses", app.createExpenseHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/expenses/:id", app.showExpenseHandler)

	return router
}
