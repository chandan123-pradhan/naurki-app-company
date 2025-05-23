package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

func AddSubscriptionPlans(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req models.AddPlansModel
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	// Validate Subscription plan params
	isValid, validationError := utils.ValidateAddSubscriptionPlan(req)
	if !isValid {
		respondWithJSON(w, http.StatusBadRequest, validationError, map[string]interface{}{})
		return
	}
	id, err := services.AddPlans(&req)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to add Plan", map[string]interface{}{})
		return
	}
	respondWithJSON(
		w, http.StatusOK, "Subscripiton plan added successfully", map[string]interface{}{
			"plan ID": id,
		},
	)

}

func GetSubscriptionPlanForAdmin(w http.ResponseWriter, r *http.Request) {

	plans, err := services.GetSubscriptionPlan()
	if err != nil {
		http.Error(w, "Failed to fetch subscription plans: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	respondWithJSON(
		w, http.StatusOK, "Subscripiton plan fetched successfully", map[string]interface{}{
			"subscription_plans": plans,
		},
	)

}

func GetSubscriptionPlanForCompanies(w http.ResponseWriter, r *http.Request) {
	// Extract JWT Token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", map[string]interface{}{})
		return
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", map[string]interface{}{})
		return
	}
	// Validate the token and extract the user ID
	_, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}
	fmt.Println("token verified not going to fetch all plans list")
	GetSubscriptionPlanForAdmin(w, r)

}




func SubscribePlan(w http.ResponseWriter, r *http.Request) {

	var req models.PlanSubscriptionModel
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", map[string]interface{}{})
		return
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", map[string]interface{}{})
		return
	}
	// Validate the token and extract the user ID
	companyId, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

	req.CompanyId = companyId
	isValid, validationError := utils.PlanSubscriptionValidate(req)
	if !isValid {
		respondWithJSON(w, http.StatusBadRequest, validationError, map[string]interface{}{})
		return
	}
	id, err := services.SubscribePlan(&req)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to Subscribe Plan", map[string]interface{}{})
		return
	}
	respondWithJSON(
		w, http.StatusOK, "Subscripiton plan added successfully", map[string]interface{}{
			"Subscription_ID": id,
		},
	)

}





func SubscriptionPayment(w http.ResponseWriter, r *http.Request) {

	var req models.SubscriptionPayment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", map[string]interface{}{})
		return
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", map[string]interface{}{})
		return
	}
	// Validate the token and extract the user ID
	companyId, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

	req.CompanyId = companyId
	isValid, validationError := utils.ValidateSubscriptionPayment(req)
	if !isValid {
		respondWithJSON(w, http.StatusBadRequest, validationError, map[string]interface{}{})
		return
	}
	id, err := services.UpdatePaymentStatus(&req)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to add payment status", map[string]interface{}{})
		return
	}
	respondWithJSON(
		w, http.StatusOK, "Subscripiton plan added successfully", map[string]interface{}{
			"Payment_id": id,
		},
	)

}



func CheckSubscriptionStatus(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", map[string]interface{}{})
		return
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", map[string]interface{}{})
		return
	}
	// Validate the token and extract the user ID
	companyId, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

status, err := services.CheckSubscriptionPlanStatus(companyId)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to fetch status", map[string]interface{}{})
		return
	}
	respondWithJSON(
		w, http.StatusOK, "Status fetched", map[string]interface{}{
			"status": status,
		},
	)

}



