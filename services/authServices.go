package services

import (
	"fmt"
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
	"naurki_app_backend.com/utils"
)

// RegisterUser handles company registration, including optional fields and profile image URL
func RegisterUser(
	companyName, companyEmail, gst, password, mobileNumber, linkedLink, websiteLink,
	address, industry, status, profileImageURL, numberOfEmployee, about string,
) (*models.Company, error) {

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create a new company instance with the provided details
	company := &models.Company{
		CompanyName:     companyName,
		CompanyEmail:    companyEmail,
		Password:        hashedPassword,
		MobileNumber:    mobileNumber,
		CompanyLogo:     profileImageURL,
		CompanyGst:      gst,
		CompanyLinkedin: linkedLink,
		CompanyWebsite:  websiteLink,
		CompanyAddress:  address,
		CompanyIndustry: industry,
		CompanyStatus:   status,
		NumberOfEmployee: numberOfEmployee,
		About:            about,
	}

	// Save the company to the database (including profile image URL)
	if err := repositories.CreateUser(company); err != nil {
		return nil, fmt.Errorf("failed to create company: %v", err)
	}

	return company, nil
}

// LoginUser handles company login, verifies credentials, and returns a company and JWT token
func LoginUser(companyEmail, password string) (*models.Company, string, error) {
	// Validate email format
	if !utils.IsValidEmail(companyEmail) {
		return nil, "", fmt.Errorf("invalid email format")
	}

	// Fetch the company from the database by email ID
	company, err := repositories.GetCompanyByEmail(companyEmail)
	if err != nil {
		return nil, "", fmt.Errorf("company not found")
	}

	// Check if the provided password matches the stored password hash
	if valid := utils.CheckPasswordHash(password, company.Password); !valid {
		return nil, "", fmt.Errorf("incorrect password")
	}

	// Generate JWT token for the company
	token, err := utils.GenerateJWT(company.ID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	// Return the sanitized company data (without the password) and the JWT token
	return &models.Company{
		ID:                   company.ID,
		CompanyName:          company.CompanyName,
		CompanyEmail:         company.CompanyEmail,
		MobileNumber:         company.MobileNumber,
		CompanyLogo:          company.CompanyLogo,
		CompanyGst:           company.CompanyGst,
		CompanyLinkedin:      company.CompanyLinkedin,
		CompanyWebsite:       company.CompanyWebsite,
		CompanyAddress:       company.CompanyAddress,
		CompanyIndustry:      company.CompanyIndustry,
		CompanyStatus:        company.CompanyStatus,
		NumberOfEmployee:     company.NumberOfEmployee,
		About:                company.About,
		CreatedAt:            company.CreatedAt,
		UpdatedAt:            company.UpdatedAt,
	}, token, nil
}
