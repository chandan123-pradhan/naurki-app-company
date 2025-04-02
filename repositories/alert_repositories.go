package repositories

import (
	"fmt"
	"time"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

func GetUserAlertsLast30Days() ([]models.UserAlert, error) {
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	stmt := `
			SELECT id, jobTitle, skills, email, username, mobile, userId, location, description, profile_image_url, created_at
			FROM user_alerts
			WHERE created_at >= ?
			ORDER BY created_at DESC
	`

	rows, err := config.DB.Query(stmt, thirtyDaysAgo)
	if err != nil {
			return nil, fmt.Errorf("could not query user alerts: %v", err)
	}
	defer rows.Close()

	var userAlerts []models.UserAlert
	for rows.Next() {
			var alert models.UserAlert
			if err := rows.Scan(
					&alert.ID, &alert.JobTitle, &alert.Skills, &alert.Email, &alert.Username,
					&alert.Mobile, &alert.UserID, &alert.Location, &alert.Description,
					&alert.ProfileImageURL, &alert.CreatedAt,
			); err != nil {
					return nil, fmt.Errorf("could not scan user alert: %v", err)
			}
			userAlerts = append(userAlerts, alert)
	}

	if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error iterating over user alerts: %v", err)
	}

	return userAlerts, nil
}