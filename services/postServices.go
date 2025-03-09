package services

import (
	"fmt"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)

// AddJobPost adds a new job post for the user and returns the job ID
func AddJobPost(userID int, jobPost struct {
	JobTitle        string `json:"job_title"`
	JobDescription  string `json:"job_description"`
	Qualification   string `json:"qualification"`
	NoOfRequirements string `json:"no_of_requirements"`
	Skills          string `json:"skills"`
	Status          string `json:"status"`
	ContactEmail    string `json:"contact_email"`
	ContactNumber   string `json:"contact_number"`
	JobLocation     string `json:"job_location"`
	CompanyLogo		 string `json:"company_logo"`
	CompanyName 	 string `json:"company_name"`
}) (int64, error) {
	// Call the repository to insert the job post and retrieve the job ID
	jobID, err := repositories.AddJobPost(
		userID,
		jobPost.JobTitle,
		jobPost.JobDescription,
		jobPost.Qualification,
		jobPost.NoOfRequirements,
		jobPost.Skills,
		jobPost.Status,
		jobPost.ContactEmail,
		jobPost.ContactNumber,
		jobPost.JobLocation,
		jobPost.CompanyLogo,
		jobPost.CompanyName,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to add job post: %v", err)
	}
	

	return jobID, nil
}


func GetJobsByCompanyID(companyID int) ([]models.Jobs, error) {
	return repositories.GetJobsByCompanyID(companyID)
}

// GetJobDetails fetches a job's details and applied users
func GetJobDetails(companyID, jobID int) (*models.JobDetails, []models.AppliedUser, error) {
	return repositories.GetJobDetailsWithApplicants(companyID, jobID)
}


func GetCompanyNotifications(companyID int) ([]models.JobApplicationLog, error) {
	return repositories.GetNotifications(companyID)
}