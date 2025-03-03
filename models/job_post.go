package models

type JobPost struct {
	JobTitle         string `json:"job_title"`
	JobDescription   string `json:"job_description"`
	Qualification    string `json:"qualification"`
	NoOfRequirements string    `json:"no_of_requirements"`
	Skills           string `json:"skills"`
	Status           string `json:"status"`
	ContactEmail     string `json:"contact_email"`
	ContactNumber    string `json:"contact_number"`
	JobLocation      string `json:"job_location"`
	CompanyLogo		 string `json:"company_logo"`
	CompanyName 	 string `json:"company_name"`
}

type Jobs struct {
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
	CreatedAt       string    `json:"created_at"`  // Change time.Time to string
}

