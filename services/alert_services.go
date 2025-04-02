package services

import (
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)


func GetAlerts() ([]models.UserAlert, error) {
	return repositories.GetUserAlertsLast30Days()
}