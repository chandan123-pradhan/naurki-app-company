package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"naurki_app_backend.com/controllers"
)

func InitializeRoutes() *mux.Router {
	// Initialize the router
	router := mux.NewRouter()

	// Authentication Routes
	router.HandleFunc("/company/register", controllers.Register).Methods("POST")
	router.HandleFunc("/company/login", controllers.Login).Methods("POST")
	router.HandleFunc("/company/get_profile",controllers.GetUserDetails).Methods("GET")
	router.HandleFunc("/company/add_post", controllers.PostNewJob).Methods("POST")
	router.HandleFunc("/company/get_post",controllers.GetCompanyJobs).Methods("GET")
	router.HandleFunc("/company/job_details", controllers.GetJobDetailsHandler).Methods("GET")
	router.HandleFunc("/company/get-notifications",controllers.GetNotificationsHandler).Methods("GET")
	router.HandleFunc("/company/get-alerts",controllers.GetAlerts).Methods("GET")
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))


	// Add other routes for authentication or any other resources
	// For example:
	// router.HandleFunc("/login", controllers.Login).Methods("POST")

	// You can add more routes here as your application grows
	// Example: router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	// Example: router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	return router
}
