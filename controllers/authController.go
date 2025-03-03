package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

// Register handles company registration, including logo upload and token generation
func Register(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10MB limit for file size
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Error parsing request", map[string]interface{}{})
		return
	}

	// Extract company information from form data
	var req struct {
		CompanyName      string `json:"company_name"`
		CompanyEmail     string `json:"company_email"`
		About            string `json:"about"`
		Password         string `json:"password"`
		ConfirmPassword  string `json:"confirm_password"`
		MobileNumber     string `json:"mobile_number"`
		CompanyGst       string `json:"gstin"`
		CompanyLinkedin  string `json:"linkedin_link"`
		CompanyWebsite   string `json:"website_link"`
		CompanyAddress   string `json:"address"`
		CompanyIndustry  string `json:"industry"`
		CompanyStatus    string `json:"status"`
		NumberOfEmployee string `json:"number_of_employee"`
	}

	// Retrieve fields from multipart form
	req.CompanyName = r.FormValue("company_name")
	req.CompanyEmail = r.FormValue("company_email")
	req.About = r.FormValue("about")
	req.Password = r.FormValue("password")
	req.ConfirmPassword = r.FormValue("confirm_password")
	req.MobileNumber = r.FormValue("mobile_number")
	req.CompanyGst = r.FormValue("gstin")
	req.CompanyLinkedin = r.FormValue("linkedin_link")
	req.CompanyWebsite = r.FormValue("website_link")
	req.CompanyAddress = r.FormValue("address")
	req.CompanyIndustry = r.FormValue("industry")
	req.CompanyStatus = r.FormValue("status")
	req.NumberOfEmployee = r.FormValue("number_of_employee")

	// Validate input
	if err := utils.ValidateCompanyRegistrationInput(req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error(), map[string]interface{}{})
		return
	}

	// Check for company logo in the form
	var logoURL string
	file, _, err := r.FormFile("company_logo")
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Company logo is required", map[string]interface{}{})
		return
	}
	defer file.Close()

	// Save the company logo to the server (or use cloud storage)
	// Generate a unique file name based on timestamp
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), "company_logo.jpg")
	filePath := filepath.Join("./uploads", fileName)

	// Create the file on the server
	outFile, err := os.Create(filePath)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Error saving file", map[string]interface{}{})
		return
	}
	defer outFile.Close()

	// Copy the content from the uploaded file to the server file
	_, err = outFile.ReadFrom(file)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Error writing file", map[string]interface{}{})
		return
	}

	// Now that the file is uploaded, save the file path (URL) to the database
	logoURL = "/uploads/" + fileName // Adjust this based on where your images are served

	// Call service to register the company with logo URL
	company, err := services.RegisterUser(req.CompanyName, 
	req.CompanyEmail, 
	req.CompanyGst, 
	req.Password, 
	req.MobileNumber, 
	req.CompanyLinkedin, 
	req.CompanyWebsite, 
	req.CompanyAddress, 
	req.CompanyIndustry, 
	req.CompanyStatus, 
	logoURL, 
	req.NumberOfEmployee, 
	req.About)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Email ID already Exists, try to login please", map[string]interface{}{})
		return
	}

	// Generate JWT token for the company
	token, err := utils.GenerateJWT(company.ID)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to generate token", map[string]interface{}{})
		return
	}

	// Prepare a sanitized company response (without password)
	companyResponse := map[string]interface{}{
		"id":                   company.ID,
		"company_name":         company.CompanyName,
		"company_email":        company.CompanyEmail,
		"about":                company.About,
		"mobile_number":        company.MobileNumber,
		"company_logo_url":     company.CompanyLogo,
		"company_gstin":        company.CompanyGst,
		"company_linkedin":     company.CompanyLinkedin,
		"company_website":      company.CompanyWebsite,
		"company_address":      company.CompanyAddress,
		"company_industry":     company.CompanyIndustry,
		"company_status":       company.CompanyStatus,
		"number_of_employee":   company.NumberOfEmployee,
	}

	// Respond with success and include the JWT token
	respondWithJSON(w, http.StatusCreated, "Registration successful", map[string]interface{}{
		"company": companyResponse, // Company data without password
		"token":   token,            // The JWT token
	})
}



// Login handles company login, authenticates using email and password, and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming request to get login credentials
	var req struct {
		CompanyEmail string `json:"email"`     // Company email
		Password     string `json:"password"`  // Company password
	}

	// Decode the incoming JSON request into the req object
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	// Validate the input fields (make sure both email and password are provided)
	if req.CompanyEmail == "" || req.Password == "" {
		respondWithJSON(w, http.StatusBadRequest, "Email and Password are required", map[string]interface{}{})
		return
	}

	// Call service to authenticate the company
	company, token, err := services.LoginUser(req.CompanyEmail, req.Password)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

	// Prepare the response company data (without password)
	companyResponse := map[string]interface{}{
		"id":                   company.ID,
		"company_name":         company.CompanyName,
		"company_email":        company.CompanyEmail,
		"mobile_number":        company.MobileNumber,
		"company_logo":         company.CompanyLogo,
		"company_gst":          company.CompanyGst,
		"company_linkedin":     company.CompanyLinkedin,
		"company_website":      company.CompanyWebsite,
		"company_address":      company.CompanyAddress,
		"company_industry":     company.CompanyIndustry,
		"company_status":       company.CompanyStatus,
		"number_of_employee":   company.NumberOfEmployee,
		"about":                company.About,
		"created_at":           company.CreatedAt,
		"updated_at":           company.UpdatedAt,
	}

	// Respond with success and include the JWT token
	respondWithJSON(w, http.StatusOK, "Login successful", map[string]interface{}{
		"company": companyResponse, // Company data without password
		"token":   token,            // The JWT token
	})
}

// Helper function to respond with JSON in the desired structure
func respondWithJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Set the response structure
	response := map[string]interface{}{
		"status":  "failure", // Default status is failure
		"message": message,
		"data":    data, // Send the provided data (or empty if nil)
	}

	// If there's no error, set the status to "success"
	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		response["status"] = "success"
	}

	// Encode the response as JSON
	json.NewEncoder(w).Encode(response)
}
