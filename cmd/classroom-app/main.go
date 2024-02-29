package main

import (
	"FinalProject/internal/classroom-app/entity"
	"database/sql"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type config struct {
	port string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models entity.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8080", "Server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "user=postgres dbname=classroom_app password=s123 host=localhost sslmode=disable", "Postgres data source")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: entity.NewModels(db),
	}
	app.run()
}

func (app *application) run() {
	r := mux.NewRouter()

	// Create class
	r.HandleFunc("/class", app.createClassHandler).Methods("POST")
	// Get class
	r.HandleFunc("/class/{classId}", app.getClassHandler).Methods("GET")
	// Update class
	r.HandleFunc("/class/{classId}", app.updateClassHandler).Methods("PUT")
	// Delete class
	r.HandleFunc("/class/{classId}", app.deleteClassHandler).Methods("DELETE")

	// Create Task
	r.HandleFunc("/task", app.createTaskHandler).Methods("POST")
	// Get Task
	r.HandleFunc("/task/{taskId}", app.getTaskHandler).Methods("GET")
	// Update Task
	r.HandleFunc("/task/{taskId}", app.updateTaskHandler).Methods("PUT")
	// Delete Task
	r.HandleFunc("/task/{taskId}", app.deleteTaskHandler).Methods("DELETE")

	// Create User
	r.HandleFunc("/user", app.createUserHandler).Methods("POST")
	// Get User
	r.HandleFunc("/user/{userId}", app.getUserHandler).Methods("GET")
	// Update User
	r.HandleFunc("/user/{userId}", app.updateUserHandler).Methods("PUT")
	// Delete User
	r.HandleFunc("/user/{userId}", app.deleteUserHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
