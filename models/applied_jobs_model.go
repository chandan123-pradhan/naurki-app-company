package models

type AppliedUser struct {
	UserID          int    `json:"user_id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	MobileNumber    string `json:"mobile_number"`
	ProfileImageURL string `json:"profile_image_url"`
	Status          string `json:"status"`
	ApplicationDate string `json:"application_date"`
}


type JobDetails struct {
	JobID           int       `json:"job_id"`
	JobTitle        string    `json:"job_title"`
	JobDescription  string    `json:"job_description"`
	Qualification   string    `json:"qualification"`
	NoOfRequirements int      `json:"no_of_requirements"`
	Skills          string    `json:"skills"`
	Status          string    `json:"status"`
	ContactEmail    string    `json:"contact_email"`
	ContactNumber   string    `json:"contact_number"`
	JobLocation     string    `json:"job_location"`
	CompanyLogo     string    `json:"company_logo"`
	CompanyName     string    `json:"company_name"`
	CreatedAt       string    `json:"created_at"`
}