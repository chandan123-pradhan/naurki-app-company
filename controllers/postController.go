package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

// PostNewJob handles the job post request
func PostNewJob(w http.ResponseWriter, r *http.Request) {
	// Extract the Authorization header
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

	// Validate the token
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

	// Parse the request body
	var req models.JobPost
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	// Validate job post
	isValid, validationError := utils.ValidateJobPost(req)
	if !isValid {
		respondWithJSON(w, http.StatusBadRequest, validationError, map[string]interface{}{})
		return
	}

	// Call the service to add the job post
	jobID, err := services.AddJobPost(userID, req)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to add job post", map[string]interface{}{})
		return
	}

	// Respond with success and the created job ID
	respondWithJSON(w, http.StatusOK, "Job posted successfully", map[string]interface{}{
		"company_id": userID,
		"job_post": map[string]interface{}{
			"job_id":             jobID,
			"job_title":          req.JobTitle,
			"job_description":    req.JobDescription,
			"qualification":      req.Qualification,
			"no_of_requirements": req.NoOfRequirements,
			"skills":             req.Skills,
			"status":             req.Status,
			"contact_email":      req.ContactEmail,
			"contact_number":     req.ContactNumber,
			"job_location":       req.JobLocation,
			"company_logo":		  req.CompanyLogo,
			"company_name":       req.CompanyName,
		},
	})
}

func GetCompanyJobs(w http.ResponseWriter, r *http.Request) {
	// Extract the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", nil)
		return
	}

	// Extract the Bearer token
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

	// Fetch job posts for the company
	jobPosts, err := services.GetJobsByCompanyID(companyID)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to fetch posted jobs", nil)
		return
	}

	// Respond with job posts
	respondWithJSON(w, http.StatusOK, "Jobs fetched successfully", map[string]interface{}{
		"company_id": companyID,
		"job_posts":  jobPosts,
	})
}



// GetJobDetailsHandler fetches job details and applied users for a specific job
func GetJobDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract and validate Authorization token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", nil)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	companyID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// Extract job ID from query parameters
	jobIDStr := r.URL.Query().Get("job_id")
	if jobIDStr == "" {
		respondWithJSON(w, http.StatusBadRequest, "Job ID is required",  map[string]interface{}{})
		return
	}

	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid Job ID",  map[string]interface{}{})
		return
	}

	// Fetch job details and applied users
	job, appliedUsers, err := services.GetJobDetails(companyID, jobID)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, err.Error(),  map[string]interface{}{})
		return
	}

	// Success response
	respondWithJSON(w, http.StatusOK, "Job details fetched successfully", map[string]interface{}{
		"job":           job,
		"applied_users": appliedUsers,
	})

}

