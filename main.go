package main

import (
	"assignment-golang-backend/db"
	"assignment-golang-backend/server"
	"log"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}
	server.Init()
}
