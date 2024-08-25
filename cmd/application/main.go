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
		log.Println("Error init database: ", err)
		return
	}

	// Init server
	e := echo.New()

	// Run server

	handler.Run(e, db)
}
