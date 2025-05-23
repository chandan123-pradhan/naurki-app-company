package repositories

import (
	"database/sql"
	"fmt"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

func AddJobPost(
	userID int, // This will act as company ID as well
	jobTitle, jobDescription, qualification string,
	noOfRequirements string, skills, status, contactEmail, contactNumber,jobLocation string, companyLogo string, companyName string,
) (int64, error) {
	stmt := `
	INSERT INTO job_post (
		jobTitle, jobDescription, qualification, noOfRequirements, contactEmail,
		contactNumber, skills, status, jobLocation, company_id, company_logo, company_name
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
	result, err := config.DB.Exec(stmt, jobTitle, jobDescription, qualification, noOfRequirements,
		contactEmail, contactNumber, skills, status, jobLocation, userID, companyLogo, companyName) // Using userID as company_id
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("could not insert job post: %v", err)
	}

	jobID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve job ID: %v", err)
	}

	return jobID, nil
}
// GetJobsByCompanyID fetches all job posts by a specific company
func GetJobsByCompanyID(companyID int) ([]models.Jobs, error) {
	stmt := `SELECT id, jobTitle, jobDescription, qualification, noOfRequirements, 
			 contactEmail, contactNumber, jobLocation, skills, status, 
			 company_name, company_logo, created_at 
			 FROM job_post WHERE company_id = ? ORDER BY created_at DESC`

	rows, err := config.DB.Query(stmt, companyID)
	if err != nil {
		fmt.Printf("Error querying posted jobs: %v\n", err)
		return nil, fmt.Errorf("could not query posted jobs: %v", err)
	}
	defer rows.Close()

	var jobPosts []models.Jobs
	for rows.Next() {
		var jobPost models.Jobs
		if err := rows.Scan(
			&jobPost.JobID, &jobPost.JobTitle, &jobPost.JobDescription, &jobPost.Qualification, 
			&jobPost.NoOfRequirements, &jobPost.ContactEmail, &jobPost.ContactNumber, 
			&jobPost.JobLocation, &jobPost.Skills, &jobPost.Status, 
			&jobPost.CompanyName, &jobPost.CompanyLogo, &jobPost.CreatedAt); err != nil {
			fmt.Printf("Error scanning job post: %v\n", err)
			return nil, fmt.Errorf("could not scan posted job: %v", err)
		}
		jobPosts = append(jobPosts, jobPost)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating over posted jobs: %v\n", err)
		return nil, fmt.Errorf("error occurred while fetching posted jobs: %v", err)
	}

	return jobPosts, nil
}


// GetJobDetailsWithApplicants fetches job details and applied users
func GetJobDetailsWithApplicants(companyID, jobID int) (*models.JobDetails, []models.AppliedUser, error) {
	var job models.JobDetails

	// SQL query to get job details ensuring it's posted by the requesting company
	jobQuery := `
		SELECT id, jobTitle, jobDescription, qualification, noOfRequirements, skills, status, 
		       contactEmail, contactNumber, jobLocation, company_logo, company_name, created_at
		FROM job_post 
		WHERE id = ? AND company_id = ?
	`

	err := config.DB.QueryRow(jobQuery, jobID, companyID).Scan(
		&job.JobID, &job.JobTitle, &job.JobDescription, &job.Qualification,
		&job.NoOfRequirements, &job.Skills, &job.Status, &job.ContactEmail,
		&job.ContactNumber, &job.JobLocation, &job.CompanyLogo,
		&job.CompanyName, &job.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil, fmt.Errorf("job not found or unauthorized")
	} else if err != nil {
		return nil, nil, fmt.Errorf("error fetching job details: %v", err)
	}

	// SQL query to get applied users
	appliedUsersQuery := `
		SELECT u.id, u.full_name, u.email_id, u.mobile_number, u.profile_image_url, a.status, a.application_date 
		FROM applications a
		JOIN users u ON a.user_id = u.id
		WHERE a.job_id = ?
	`

	rows, err := config.DB.Query(appliedUsersQuery, jobID)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching applied users: %v", err)
	}
	defer rows.Close()

	appliedUsers := []models.AppliedUser{}
	for rows.Next() {
		var user models.AppliedUser
		if err := rows.Scan(&user.UserID, &user.FullName, &user.Email, &user.MobileNumber,
			&user.ProfileImageURL, &user.Status, &user.ApplicationDate); err != nil {
			return nil, nil, fmt.Errorf("error scanning applied user: %v", err)
		}
		appliedUsers = append(appliedUsers, user)
	}

	return &job, appliedUsers, nil
}





// GetNotifications fetches job application logs for a specific company
func GetNotifications(companyID int) ([]models.JobApplicationLog, error) {
	var logs []models.JobApplicationLog

	query := `
		SELECT id, user_id, job_id, company_id, job_title, user_name, profile_pic, applied_at
		FROM job_applications_log 
		WHERE company_id = ?
		ORDER BY applied_at DESC
	`

	rows, err := config.DB.Query(query, companyID)
	if err != nil {
		return nil, fmt.Errorf("error fetching notifications: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var log models.JobApplicationLog
		if err := rows.Scan(&log.ID, &log.UserID, &log.JobID, &log.CompanyID, &log.JobTitle, &log.UserName, &log.ProfilePic, &log.AppliedAt); err != nil {
			return nil, fmt.Errorf("error scanning notification: %v", err)
		}
		logs = append(logs, log)
	}
	if(len(logs)==0){
		return []models.JobApplicationLog{}, nil
	}

	return logs, nil
}


