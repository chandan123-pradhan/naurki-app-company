package services

import (
	"fmt"
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)


// GetUserDetails fetches the user's details along with employment history
func GetUserDetails(userID int) (*models.CompanyModel, error) {
	// Fetch user details from the database (this would include basic user info like name, email, etc.)
	company, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user details: %v", err)
	}

	
	// Return the user details along with employment history
	return &models.CompanyModel{
		ID:                 company.ID,
		CompanyName: company.CompanyName,
		CompanyEmail:company.CompanyEmail,
		About:company.About,
		MobileNumber:company.MobileNumber,
		CompanyLogo:company.CompanyLogo,
		CompanyGst:company.CompanyGst,
		CompanyLinkedin:company.CompanyLinkedin,
		CompanyWebsite:company.CompanyWebsite,
		CompanyAddress:company.CompanyAddress,
		CompanyIndustry:company.CompanyIndustry,
		CompanyStatus:company.CompanyStatus,
		NumberOfEmployee:company.NumberOfEmployee,
		CreatedAt:company.CreatedAt,
		UpdatedAt:company.UpdatedAt,
	}, nil
}
