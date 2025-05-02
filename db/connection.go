package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	driveNameDbConst string = "postgres"
	hostDbConst      string = "localhost"
	portDbConst      int    = 5432
	userDbConst      string = "admin"
	passwordDbConst  string = "1234"
	dbNameConst      string = "dbapp"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open(driveNameDbConst, defineEnvConection())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("erro ao configurar driver do migrate: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		driveNameDbConst, driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		m.Down()
		log.Fatal("erro ao aplicar migrations: ", err)
	}

	return db, err
}

func defineEnvConection() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostDbConst,
		portDbConst,
		userDbConst,
		passwordDbConst,
		dbNameConst,
	)
}
