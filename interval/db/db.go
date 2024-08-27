package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//
//	connect := "user=postgres password=durka dbname=TaskTracker host=host.docker.internal port=5432 sslmode=disable"
//

func InitDb() (*sql.DB, error) {
	connect := "user=postgres password=durka dbname=TaskTracker host=host.docker.internal port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
