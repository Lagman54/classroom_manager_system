package main

import (
	"FinalProject/internal/classroom-app/model"
	"FinalProject/internal/classroom-app/validator"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Header       string `json:"header"`
		Description  string `json:"description"`
		ClassroomIds []int  `json:"classrooms"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	task := &model.Task{
		Header:      input.Header,
		Description: input.Description,
	}

	v := validator.New()
	if model.ValidateTask(v, task); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Tasks.Insert(task, input.ClassroomIds)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := app.readIDParam(r)
	if err != nil || taskId < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	task, err := app.models.Tasks.Get(taskId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Task with ID %d not found", taskId)
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
}

func (app *application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := app.readIDParam(r)
	if err != nil || taskId < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	task, err := app.models.Tasks.Get(taskId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Task with ID %d not found", taskId)
		}
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Header      *string `json:"header"`
		Description *string `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Header != nil {
		task.Header = *input.Header
	}
	if input.Description != nil {
		task.Description = *input.Description
	}

	v := validator.New()
	if model.ValidateTask(v, task); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Tasks.Update(task)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
}

func (app *application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := app.readIDParam(r)
	if err != nil || taskId < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Tasks.Delete(taskId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.writeJSON(w, http.StatusOK, envelope{"result": "Success"}, nil)
}
