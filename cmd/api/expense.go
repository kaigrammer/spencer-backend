package main

import (
	"fmt"
	"net/http"
)

func (app *application) createExpenseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new expense")
}

func (app *application) showExpenseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of expense %d\n", id)
}
