package repositories

import (
	"database/sql"
	"fmt"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

// CreateUser inserts a new user into the database
func CreateUser(company *models.Company) error {
	// Prepare SQL statement (using ? placeholders for MySQL)
	stmt := `INSERT INTO companies (companyName, email, about, password, mobile_number, logo, gstin, linkedProfileLink, websiteLink, address, industry, status, number_of_employees) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ? ,?, ? ,? , ?, ?)`

	// Execute the query
	result, err := config.DB.Exec(stmt, company.CompanyName, company.CompanyEmail, company.About, company.Password, company.MobileNumber, company.CompanyLogo, company.CompanyGst, company.CompanyLinkedin, company.CompanyWebsite, company.CompanyAddress, company.CompanyIndustry, company.CompanyStatus, company.NumberOfEmployee)
	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	// Get the last inserted ID (auto_increment value)
	userID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not retrieve last inserted ID: %v", err)
	}

	// Set the user ID to the newly inserted ID
	company.ID = int(userID)

	return nil
}
func GetCompanyByEmail(emailID string) (*models.Company, error) {
	// Prepare the SQL query to fetch company details by email ID
	stmt := `SELECT id, companyName, email, about, password, mobile_number, logo, gstin, linkedProfileLink, websiteLink, address, industry, status, number_of_employees
			 FROM companies WHERE email = ?`

	// Execute the query and store the result
	var company models.Company
	err := config.DB.QueryRow(stmt, emailID).Scan(
		&company.ID,
		&company.CompanyName,
		&company.CompanyEmail,
		&company.About,
		&company.Password,
		&company.MobileNumber,
		&company.CompanyLogo,
		&company.CompanyGst,
		&company.CompanyLinkedin,
		&company.CompanyWebsite,
		&company.CompanyAddress,
		&company.CompanyIndustry,
		&company.CompanyStatus,
		&company.NumberOfEmployee,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("company not found")
		}
		return nil, fmt.Errorf("failed to fetch company: %v", err)
	}

	return &company, nil
}
