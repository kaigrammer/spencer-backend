package main

import (
	"log"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     VERSION,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
