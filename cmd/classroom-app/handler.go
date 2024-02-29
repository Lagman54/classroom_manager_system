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

func (app *application) getClassHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) updateClassHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) deleteClassHandler(w http.ResponseWriter, r *http.Request) {
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

// Task handlers

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Header      string `json:"header"`
		Description string `json:"description"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	task := &entity.Task{
		Header:      input.Header,
		Description: input.Description,
	}

	err = app.models.Tasks.Insert(task)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	app.respondWithJSON(w, http.StatusOK, task)
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["taskId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	task, err := app.models.Tasks.Get(taskId)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, task)
}

func (app *application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["taskId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	task, err := app.models.Tasks.Get(taskId)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Header      *string `json:"header"`
		Description *string `json:"description"`
	}

	err = app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if input.Header != nil {
		task.Header = *input.Header
	}
	if input.Description != nil {
		task.Description = *input.Description
	}

	err = app.models.Tasks.Update(task)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, task)
}

func (app *application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["taskId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	err = app.models.Tasks.Delete(taskId)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal sever error")
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "Success"})
}

// User handlers

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	task := &entity.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	err = app.models.Users.Insert(task)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, task)
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid user Id")
		return
	}

	user, err := app.models.Users.Get(userId)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, user)
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid user Id")
		return
	}

	var input struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
	}

	err = app.readJSON(r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := app.models.Users.Get(userId)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		user.LastName = *input.LastName
	}

	err = app.models.Users.Update(user)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, user)
}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid user Id")
		return
	}

	err = app.models.Users.Delete(userId)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "Success"})
}
