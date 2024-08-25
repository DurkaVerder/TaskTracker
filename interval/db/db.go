package db

import "database/sql"

func InitDb() (*sql.DB, error) {
	connect := "user=postgres password=durka dbname=taskTracker host=docker.interval.default port=5432"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
