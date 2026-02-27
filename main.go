package main

import (
	"Recall/cmd"
	"Recall/internal/db"
	"log"
)

func main() {
	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()
	
	// Run CLI
	cmd.Execute()
}