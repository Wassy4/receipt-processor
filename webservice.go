package main

import (
	"log"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	HandleRequests(db)
}
