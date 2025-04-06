package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors" // ✅ Import the CORS package

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

	// Initialize DB and Firebase
	config.InitDB()
	firebaseconfig.InitFirebase()

	// Initialize all routes
	router := routes.InitializeRoutes()

	// ✅ Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 🔒 Replace "*" with actual origins in production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler)) // 👈 Use corsHandler here
}
