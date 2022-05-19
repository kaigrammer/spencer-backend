package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "status: available, environment: %s, version: %s\n", app.config.env, VERSION)
	})

	return router
}
