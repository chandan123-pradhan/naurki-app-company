package controllers

import (
	"encoding/json"

	"net/http"
	"strings"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/services" // Import services package
	"naurki_app_backend.com/utils"
)


func UpdateFcmToken(w http.ResponseWriter, r *http.Request) {
	// Get Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", nil)
		return
	}

	// Extract Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", nil)
		return
	}

	// Validate the token and extract company ID
	companyID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// Parse the request body to extract the FCM token
	var request models.UpdateFcmTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	// Ensure the FCM token is provided
	if request.FcmToken == "" {
		respondWithJSON(w, http.StatusBadRequest, "FCM token is required", nil)
		return
	}

	// Call the service to update the FCM token
	err = services.UpdateFcmToken(companyID, request.FcmToken)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to update FCM token", nil)
		return
	}

	// Respond with success
	respondWithJSON(w, http.StatusOK, "FCM token updated successfully", map[string]interface{}{})
}


func SendNotification(w http.ResponseWriter, r *http.Request) {
	var req models.SendNotificationRequest

	// Parse the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	// Basic validation
	if req.FcmToken == "" || req.Title == "" || req.Body == "" {
		respondWithJSON(w, http.StatusBadRequest, "All fields (fcm_token, title, body) are required", nil)
		return
	}

	// Call the service to send the notification
	err := services.SendNotificationToToken(req.FcmToken, req.Title, req.Body)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to send notification", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, "Notification sent successfully", nil)
}
