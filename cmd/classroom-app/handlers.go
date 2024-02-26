package main

import (
	"FinalProject/internal/classroom-app/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *application) readJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

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
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	classroom := &entity.Classroom{
		Name:        input.Name,
		Description: input.Description,
	}

	err = app.models.Classrooms.Insert(classroom)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal server error")
		return
	}
	app.respondWithJSON(w, http.StatusOK, classroom)
}

func (app *application) createGetClassHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) createUpdateClassHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["classId"])
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	class, err := app.models.Classrooms.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 not found")
		return
	}

	var input struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
	}

	err = app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if input.Name != nil {
		class.Name = *input.Name
	}

	if input.Description != nil {
		class.Description = *input.Description
	}

	err = app.models.Classrooms.Update(class)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, class)
}

func (app *application) createDeleteClassHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["classId"])
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	err = app.models.Classrooms.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "Success"})
}
