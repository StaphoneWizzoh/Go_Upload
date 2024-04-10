package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("file_storage.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// Auto migrate the database
	err = db.AutoMigrate(&File{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	router := mux.NewRouter()
	initializeRoutes(router, db)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
