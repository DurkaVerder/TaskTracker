package main

import (
	"TaskTracker/interval/db"
	"log"
)

func main() {
	// Init database
	db, err := db.InitDb()
	if err != nil {
		log.Println("Error init database: ", err)
		return
	}

	// Init server

	// Run server
}