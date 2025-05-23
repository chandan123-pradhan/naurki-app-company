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
	router.HandleFunc("/company/update-fcm",controllers.UpdateFcmToken).Methods("POST")
	router.HandleFunc("/company/send-notification",controllers.SendNotification).Methods("POST")
    router.HandleFunc("/company/get_subscription_plan", controllers.GetSubscriptionPlanForCompanies).Methods("GET")
	
	router.HandleFunc("/company/payment_update", controllers.SubscriptionPayment).Methods("POST")
	router.HandleFunc("/company/subsribe_plan", controllers.SubscribePlan).Methods("POST")	
	router.HandleFunc("/company/check_subscription_status", controllers.CheckSubscriptionStatus).Methods("GET")		
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))




	//Admins....

	router.HandleFunc("/admin/add-plans",controllers.AddSubscriptionPlans).Methods("POST")
    router.HandleFunc("/admin/get-plans",controllers.GetSubscriptionPlanForAdmin).Methods("GET")


	// Add other routes for authentication or any other resources
	// For example:
	// router.HandleFunc("/login", controllers.Login).Methods("POST")

	// You can add more routes here as your application grows
	// Example: router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	// Example: router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	return router
}
