package main

import (
	"TaskTracker/interval/db"
	"TaskTracker/interval/handler"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Init database
	db, err := db.InitDb()
	if err != nil {
		log.Fatal("Error init database: ", err)
	}
	defer db.Close()

	// Init server
	e := echo.New()

	// Run server

	handler.Run(e, db)
}
