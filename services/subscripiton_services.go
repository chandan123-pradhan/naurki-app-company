package services

import (
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)

func AddPlans(addPlan *models.AddPlansModel) (int64, error) {
	return repositories.AddSubscriptionPlan(addPlan)
}


func GetSubscriptionPlan() ([]models.GetSubscriptionPlanModel, error) {
	plans, err := repositories.GetSubscriptionPlans()
	if err != nil {
		return nil, err
	}
	return plans, nil
}


func UpdatePaymentStatus(subscriptionPaymentModel *models.SubscriptionPayment) (int64, error) {
	return repositories.UpdatePaymentStatus(subscriptionPaymentModel)
}

func SubscribePlan(subscriptionPaymentModel *models.PlanSubscriptionModel) (int64, error) {
	return repositories.SubscribePlan(subscriptionPaymentModel)
}



func CheckSubscriptionPlanStatus(companyId int) (string, error) {
	return repositories.CheckSubscriptionPlanStatus(companyId)
}
