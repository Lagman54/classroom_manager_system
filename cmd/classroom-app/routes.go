package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)

	r.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/healthcheck", app.healthcheckHandler).Methods("GET")

	// Create class
	api.HandleFunc("/class", app.requireActivatedUser(app.createClassHandler)).Methods("POST")
	// Get class
	api.HandleFunc("/class/{id}", app.requireActivatedUser(app.getClassHandler)).Methods("GET")
	// Get list of classrooms
	api.HandleFunc("/classes", app.requireActivatedUser(app.getClassesList)).Methods("GET")
	// Update class
	api.HandleFunc("/class/{id}", app.requireActivatedUser(app.updateClassHandler)).Methods("PUT")
	// Delete class
	api.HandleFunc("/class/{id}", app.requirePermissions("class:write", app.deleteClassHandler)).Methods("DELETE")
	// Get tasks of a class
	api.HandleFunc("/class/{id}/tasks", app.requireActivatedUser(app.getTasksForClass)).Methods("GET")

	// Create Task
	api.HandleFunc("/task", app.requirePermissions("task:write", app.createTaskHandler)).Methods("POST")
	// Get Task
	api.HandleFunc("/task/{id}", app.requirePermissions("task:read", app.getTaskHandler)).Methods("GET")
	// Update Task
	api.HandleFunc("/task/{id}", app.requirePermissions("task:write", app.updateTaskHandler)).Methods("PUT")
	// Delete Task
	api.HandleFunc("/task/{id}", app.requirePermissions("task:write", app.deleteTaskHandler)).Methods("DELETE")

	// User handlers with Authentication
	api.HandleFunc("/user", app.registerUserHandler).Methods("POST")
	api.HandleFunc("/user/activated", app.activateUserHandler).Methods("PUT")
	api.HandleFunc("/user/login", app.createAuthenticationTokenHandler).Methods("POST")

	// Wrap the router with the panic recovery middleware and rate limit middleware.
	return app.authenticate(r)
}
