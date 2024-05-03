package main

import (
	"FinalProject/internal/classroom-app/model"
	"FinalProject/internal/classroom-app/model/filler"
	"database/sql"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/peterbourgon/ff/v3"
	"log"
	"os"
	"sync"
)

type config struct {
	port       int
	env        string
	fill       bool
	migrations string
	db         struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
	wg     sync.WaitGroup
}

func main() {
	fs := flag.NewFlagSet("demo-app", flag.ContinueOnError)

	var (
		cfg        config
		fill       = fs.Bool("fill", false, "Fill database with dummy data")
		migrations = fs.String("migrations", "", "Path to migration files folder. If not provided, migrations do not applied")
		port       = fs.Int("port", 8080, "API server port")
		env        = fs.String("env", "development", "Environment (development|staging|production)")
		dbDsn      = fs.String("dsn", "postgres://postgres:s123@localhost:5432/classroom_app?sslmode=disable", "PostgreSQL DSN")
	)

	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars()); err != nil {
		log.Fatal(err, nil)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	cfg.port = *port
	cfg.env = *env
	cfg.fill = *fill
	cfg.db.dsn = *dbDsn
	cfg.migrations = *migrations

	log.Println("starting application with configuration", map[string]string{
		"port":       fmt.Sprintf("%d", cfg.port),
		"fill":       fmt.Sprintf("%t", cfg.fill),
		"env":        cfg.env,
		"db":         cfg.db.dsn,
		"migrations": cfg.migrations,
	})

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal("DB open: " + err.Error())
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
	if cfg.migrations != "" {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		m, err := migrate.NewWithDatabaseInstance(
			cfg.migrations,
			"postgres", driver)
		if err != nil {
			return nil, err
		}
		err = m.Up()
		if err != nil && err.Error() != "no change" {
			return nil, err
		}
	}
	return db, nil
}
