package utils

import "naurki_app_backend.com/models"

func ValidateAddSubscriptionPlan(req models.AddPlansModel) (bool, string) {
	// Validate required fields
	if req.Name == "" {
		return false, "Plan Name is required."
	}
	if req.Description == "" {
		return false, "Plan Description is required."
	}
	if req.BillingCycle == "" {
		return false, "Billing Cycle is required."
	}
	if req.BillingCycle != "monthly" && req.BillingCycle != "yearly" {
	return false, "Billing cycle must be either 'monthly' or 'yearly'"
}	
	if req.Price ==0 {
		return false, "Price is requried."
	}
	return true, ""
}


func ValidateSubscriptionPayment(req models.SubscriptionPayment) (bool, string) {
	// Validate required fields
	if req.PaymentMethod == "" {
		return false, "Payment method required."
	}
	if req.PaymentMethod != "upi" && req.PaymentMethod != "netbanking" {
		return false, "Invalid Payment method"
	}
	if req.Status == "" {
		return false, "Status is required."
	}
	if req.TransactionId == "" {
		return false, "Transaction ID is required."
	}
	
	if req.SubscriptionId =="" {
		return false, "Subscripiton ID is requried."
	}
	return true, ""
}


func PlanSubscriptionValidate(req models.PlanSubscriptionModel) (bool, string) {
	// Validate required fields
	if req.PlanId == 0 {
		return false, "Plan ID is required"
	}
	if req.StartDate == "" || req.EndDate == "" {
		return false, "Start and End date required"
	}
	if req.Status == "" {
		return false, "Status is required."
	}
	return true, ""
}