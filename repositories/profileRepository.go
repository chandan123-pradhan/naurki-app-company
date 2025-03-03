package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)
func GetUserByID(userID int) (*models.CompanyModel, error) {
	// Prepare SQL query to fetch the user details
	stmt := `SELECT id, companyName, email, about, mobile_number, logo, gstin, linkedProfileLink, websiteLink, 
	address, industry, status, number_of_employees, created_at, updated_at
			 FROM companies WHERE id = ?`

	// Execute the query and scan the result into the user model
	var company models.CompanyModel
	err := config.DB.QueryRow(stmt, userID).Scan(
		&company.ID, 
		&company.CompanyName, 
		&company.CompanyEmail, 
		&company.About, 
		&company.MobileNumber, 
		&company.CompanyLogo, 
		&company.CompanyGst, 
		&company.CompanyLinkedin, 
		&company.CompanyWebsite, 
		&company.CompanyAddress, 
		&company.CompanyIndustry, 
		&company.CompanyStatus, 
		&company.NumberOfEmployee, 
		&company.CreatedAt, 
		&company.UpdatedAt,
	)

if err != nil {
    if err == sql.ErrNoRows {
        log.Printf("User with ID %d not found", userID)
        return nil, fmt.Errorf("user not found")
    }
    log.Printf("Error fetching user details: %v", err)  // Log the actual error
    return nil, fmt.Errorf("failed to fetch user: %v", err)
}

	return &company, nil
}
