package repositories

import (
	"fmt"
	"log"

	"naurki_app_backend.com/config"
)

func UpdateOrAddFcmToken(companyID int, fcmToken string) error {
	// First, try to update the FCM token for the company
	stmt := `
		UPDATE company_fcm_tokens
		SET fcm_token = ?, updated_at = CURRENT_TIMESTAMP
		WHERE company_id = ?
	`

	// Execute the query to update the FCM token
	result, err := config.DB.Exec(stmt, fcmToken, companyID)
	if err != nil {
		log.Printf("Error updating FCM token for company ID %d: %v", companyID, err)
		return fmt.Errorf("failed to update FCM token: %v", err)
	}

	// Check how many rows were affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting affected rows for company ID %d: %v", companyID, err)
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		// If no rows were affected, the record does not exist, so insert a new record
		insertStmt := `
			INSERT INTO company_fcm_tokens (company_id, fcm_token)
			VALUES (?, ?)
		`

		// Execute the insert query
		_, err := config.DB.Exec(insertStmt, companyID, fcmToken)
		if err != nil {
			log.Printf("Error inserting FCM token for company ID %d: %v", companyID, err)
			return fmt.Errorf("failed to insert FCM token: %v", err)
		}

		log.Printf("FCM token inserted successfully for company ID %d", companyID)
		return nil
	}

	// If rows were affected, the update was successful
	log.Printf("FCM token updated successfully for company ID %d", companyID)
	return nil
}
