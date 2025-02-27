package main

import (
	"FinalProject/internal/classroom-app/model"
	"FinalProject/internal/classroom-app/validator"
	"database/sql"
	"errors"
	"log"
	"net/http"
)

func (app *application) createClassHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	classroom := &model.Classroom{
		Name:        input.Name,
		Description: input.Description,
	}

	v := validator.New()
	if model.ValidateClassroom(v, classroom); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Classrooms.Insert(classroom)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusCreated, envelope{"classroom": classroom}, nil)
}

func (app *application) getClassHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	classroom, err := app.models.Classrooms.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Class with ID %d not found", id)
		}
		app.notFoundResponse(w, r)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"classroom": classroom}, nil)
}

func (app *application) getClassesList(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string
		model.Filters
	}
	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readStrings(qs, "name", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readStrings(qs, "sort", "id")

	input.Filters.SortSafeList = []string{
		"id", "name",
		"-id", "-name",
	}

	if model.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	classrooms, metadata, err := app.models.Classrooms.GetAll(input.Name, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"classrooms": classrooms, "metadata": metadata}, nil)
}

func (app *application) updateClassHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	classroom, err := app.models.Classrooms.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Class with ID %d not found", id)
		}
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		classroom.Name = *input.Name
	}
	if input.Description != nil {
		classroom.Description = *input.Description
	}

	v := validator.New()
	if model.ValidateClassroom(v, classroom); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Classrooms.Update(classroom)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"classroom": classroom}, nil)
}

func (app *application) deleteClassHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Classrooms.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"result": "Success"}, nil)
}
