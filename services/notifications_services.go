package services

import (
	"context"
	"fmt"
	"log"

	firebaseconfig "naurki_app_backend.com/firebase_config"
	"naurki_app_backend.com/repositories" // Import the repositories package

	firebase_messaging "firebase.google.com/go/v4/messaging"
)

// UpdateFcmToken is a service function that calls the repository to update the FCM token
func UpdateFcmToken(userID int, fcmToken string) error {
	// Call the repository function to update the FCM token
	err := repositories.UpdateOrAddFcmToken(userID, fcmToken)
	if err != nil {
		// Handle error and return it if the update failed
		return fmt.Errorf("service: failed to update FCM token: %w", err)
	}

	// If everything goes fine, return nil
	return nil
}




func SendNotificationToToken(token string, title string, body string) error {
    ctx := context.Background()

    client := firebaseconfig.GetClient() // gets the initialized Firebase client

    message := &firebase_messaging.Message{
        Token: token,
        Notification: &firebase_messaging.Notification{
            Title: title,
            Body:  body,
        },
    }

    response, err := client.Send(ctx, message)
    if err != nil {
        log.Printf("Failed to send notification: %v", err)
        return err
    }

    log.Printf("Successfully sent message: %s", response)
    return nil
}