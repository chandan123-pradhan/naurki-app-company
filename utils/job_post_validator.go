package utils

import "naurki_app_backend.com/models"

func ValidateJobPost(req models.JobPost) (bool, string) {
	// Validate required fields
	if req.JobTitle == "" {
		return false, "Job Title is required"
	}
	if req.JobDescription == "" {
		return false, "Job Description is required"
	}
	if req.Qualification == "" {
		return false, "Qualification is required"
	}
	if req.NoOfRequirements =="" {
		return false, "No of Requirements must be a positive number"
	}
	if req.Skills == "" {
		return false, "Skills cannot be empty"
	}
	if req.Status == "" {
		return false, "Status is required"
	}
	validStatuses := []string{"open", "closed", "on_hold", "filled"}
	isValidStatus := false
	for _, validStatus := range validStatuses {
		if req.Status == validStatus {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		return false, "Invalid Status value. Allowed values are: open, closed, on_hold, filled"
	}
	if req.ContactEmail != "" && !isValidEmail(req.ContactEmail) {
		return false, "Invalid contact email format"
	}
	if req.ContactNumber != "" && !isValidPhoneNumber(req.ContactNumber) {
		return false, "Invalid contact number format"
	}
	if req.JobLocation == "" {
		return false, "Job Location is required"
	}
	return true, ""
}