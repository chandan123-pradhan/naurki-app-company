package controllers

import (
	"net/http"
	"strings"

	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)





func GetAlerts(w http.ResponseWriter, r *http.Request) {
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

	// Fetch the user's details including employment history
	userDetails, err := services.GetAlerts()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error(), map[string]interface{}{})
		return
	}

	// Respond with the user's details
	respondWithJSON(w, http.StatusOK, "User details fetched successfully", map[string]interface{}{
		"user": userDetails,
	})
}