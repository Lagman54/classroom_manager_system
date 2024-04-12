package main

import (
	"FinalProject/internal/classroom-app/entity"
	"database/sql"
	"flag"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models entity.Models
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "user=postgres dbname=classroom_app password=s123 host=localhost sslmode=disable", "Postgres data source")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	app := &application{
		config: cfg,
		models: entity.NewModels(db),
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

//func (app *application) run() {
//	r := mux.NewRouter()
//
//	// Create class
//	r.HandleFunc("/class", app.createClassHandler).Methods("POST")
//	// Get class
//	r.HandleFunc("/class/{id}", app.getClassHandler).Methods("GET")
//	// Update class
//	r.HandleFunc("/class/{id}", app.updateClassHandler).Methods("PUT")
//	// Delete class
//	r.HandleFunc("/class/{id}", app.deleteClassHandler).Methods("DELETE")
//
//	// Create Task
//	r.HandleFunc("/task", app.createTaskHandler).Methods("POST")
//	// Get Task
//	r.HandleFunc("/task/{id}", app.getTaskHandler).Methods("GET")
//	// Update Task
//	r.HandleFunc("/task/{id}", app.updateTaskHandler).Methods("PUT")
//	// Delete Task
//	r.HandleFunc("/task/{id}", app.deleteTaskHandler).Methods("DELETE")
//
//	//// Create User
//	//r.HandleFunc("/user", app.createUserHandler).Methods("POST")
//	//// Get User
//	//r.HandleFunc("/user/{id}", app.getUserHandler).Methods("GET")
//	//// Update User
//	//r.HandleFunc("/user/{id}", app.updateUserHandler).Methods("PUT")
//	//// Delete User
//	//r.HandleFunc("/user/{id}", app.deleteUserHandler).Methods("DELETE")
//
//	log.Printf("Starting server on %s\n", app.config.port)
//	err := http.ListenAndServe(app.config.port, r)
//	log.Fatal(err)
//}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
