package models





// Request body structure to capture the FCM token
type UpdateFcmTokenRequest struct {
	FcmToken string `json:"fcm_token"`  // Field for the FCM token
}