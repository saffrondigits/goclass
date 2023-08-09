package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func DBConn() (*sql.DB, error) {
	connStr := "postgres://root:abcd@localhost:1234/goclass?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Errorf("cannot connect to the db: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logrus.Errorf("cannot ping to the db: %v", err)
		return nil, err
	}

	return db, nil
}
