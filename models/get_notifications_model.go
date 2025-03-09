package models


// JobApplicationLog represents an entry in the job applications log
type JobApplicationLog struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	JobID      int       `json:"job_id"`
	CompanyID  int       `json:"company_id"`
	JobTitle   string    `json:"job_title"`
	UserName   string    `json:"user_name"`
	ProfilePic string    `json:"profile_pic"`
	AppliedAt  string `json:"applied_at"`
}
