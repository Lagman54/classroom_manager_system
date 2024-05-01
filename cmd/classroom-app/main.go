package main

import (
	"FinalProject/internal/classroom-app/model"
	"FinalProject/internal/classroom-app/model/filler"
	"database/sql"
	"flag"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	fill bool
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	flag.BoolVar(&cfg.fill, "fill", false, "Fill database with dummy data")
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
		models: model.NewModels(db),
	}

	if cfg.fill {
		err := filler.PopulateDatabase(app.models)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
