package main

import (
	"FinalProject/internal/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createClassHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	classroom := &entity.Classroom{
		Name: input.Name,
	}

	err = app.models.Classrooms.Insert(classroom)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal server error")
		return
	}
	app.respondWithJSON(w, http.StatusOK, classroom)
}

func (app *application) createGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["classId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid classroom ID")
		return
	}

	classroom, err := app.models.Classrooms.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, classroom)
}

func (app *application) readJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
