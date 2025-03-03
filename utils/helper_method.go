package utils

import (
	"fmt"
	"regexp"
)

// Validate company registration input
func ValidateCompanyRegistrationInput(req struct {
	CompanyName     string `json:"company_name"`
	CompanyEmail    string `json:"company_email"`
	About           string `json:"about"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	MobileNumber    string `json:"mobile_number"`
	CompanyGst      string `json:"gstin"`
	CompanyLinkedin string `json:"linkedin_link"`
	CompanyWebsite  string `json:"website_link"`
	CompanyAddress  string `json:"address"`
	CompanyIndustry string `json:"industry"`
	CompanyStatus   string `json:"status"`
	NumberOfEmployee string `json:"number_of_employee"`
}) error {

	// Validate company name
	if req.CompanyName == "" {
		return fmt.Errorf("Company name is required")
	}
	if len(req.CompanyName) < 3 {
		return fmt.Errorf("Company name must be at least 3 characters long")
	}

	// Validate company email
	if req.CompanyEmail == "" {
		return fmt.Errorf("Company email is required")
	}
	if !isValidEmail(req.CompanyEmail) {
		return fmt.Errorf("Invalid email format")
	}

	// Validate about section
	if req.About == "" {
		return fmt.Errorf("About the company is required")
	}
	if len(req.About) < 10 {
		return fmt.Errorf("About section must be at least 10 characters long")
	}

	// Validate password
	if req.Password == "" {
		return fmt.Errorf("Password is required")
	}
	if len(req.Password) < 8 {
		return fmt.Errorf("Password must be at least 8 characters long")
	}

	// Validate password confirmation
	if req.Password != req.ConfirmPassword {
		return fmt.Errorf("Passwords do not match")
	}

	// Validate mobile number
	if req.MobileNumber == "" {
		return fmt.Errorf("Mobile number is required")
	}
	if len(req.MobileNumber) != 10 {
		return fmt.Errorf("Mobile number must be exactly 10 digits")
	}

	// Validate GSTIN
	if req.CompanyGst != "" && !isValidGSTIN(req.CompanyGst) {
		return fmt.Errorf("Invalid GSTIN format")
	}

	// Validate LinkedIn URL (optional)
	if req.CompanyLinkedin != "" && !isValidURL(req.CompanyLinkedin) {
		return fmt.Errorf("Invalid LinkedIn URL")
	}

	// Validate company website (optional)
	if req.CompanyWebsite != "" && !isValidURL(req.CompanyWebsite) {
		return fmt.Errorf("Invalid website URL")
	}

	// Validate company address
	if req.CompanyAddress == "" {
		return fmt.Errorf("Company address is required")
	}

	// Validate company industry
	if req.CompanyIndustry == "" {
		return fmt.Errorf("Industry is required")
	}

	// Validate company status
	if req.CompanyStatus == "" {
		return fmt.Errorf("Company status is required")
	}

	// Validate number of employees
	if req.NumberOfEmployee == "" {
		return fmt.Errorf("Number of employees is required")
	}
	

	return nil
}

// Helper function to validate email
func isValidEmail(email string) bool {
	// Basic regex for email validation
	// You can use a more strict regex based on your needs
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Helper function to validate GSTIN (example format: "AAAPL1234C1Z5")
func isValidGSTIN(gstin string) bool {
	if len(gstin) != 15 {
		return false
	}
	re := regexp.MustCompile(`^[A-Z]{2}[0-9]{4}[A-Z]{5}[0-9]{4}[A-Z]{1}[1-9A-Z]{1}[A-Z0-9]{1}$`)
	return re.MatchString(gstin)
}

// Helper function to validate URLs (LinkedIn and Website)
func isValidURL(url string) bool {
	// Simple URL validation using regex for basic format
	re := regexp.MustCompile(`^(https?://)?(www\.)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}(/[\w-]*)*$`)
	return re.MatchString(url)
}
