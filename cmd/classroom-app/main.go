package main

import (
	"FinalProject/internal/entity"
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

	r.HandleFunc("/class", app.createClassHandler).Methods("POST")
	r.HandleFunc("/class/{classId}", app.createGetHandler).Methods("GET")

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
