package models

// UserDetailsResponse represents the user's details along with their employment history
type UserDetailsResponse struct {
	ID                  int                    `json:"id"`
	FullName            string                 `json:"full_name"`
	EmailID             string                 `json:"email_id"`
	HighestQualification string                `json:"highest_qualification"`
	MobileNumber        string                 `json:"mobile_number"`
	ProfileImageURL     string                 `json:"profile_image_url"`
	EmploymentHistory   []EmploymentDetail     `json:"employment_history"`
}

// EmploymentDetail represents an employment record
type EmploymentDetail struct {
	EmployerName string `json:"employer_name"`
	Designation  string `json:"designation"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}
