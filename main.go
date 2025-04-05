package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"naurki_app_backend.com/config"
	firebaseconfig "naurki_app_backend.com/firebase_config"
	"naurki_app_backend.com/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection (example function in config package)
	config.InitDB()

	firebaseconfig.InitFirebase() 

	// Initialize all routes
	router := routes.InitializeRoutes()

	// Start the HTTP server
	port := os.Getenv("PORT") // Use the PORT value from .env
	if port == "" {
		port = "8080" // Default to 8080 if no PORT is set in .env
	}
	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
