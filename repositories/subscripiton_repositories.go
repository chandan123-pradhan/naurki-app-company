package repositories

import (
	"database/sql"
	"fmt"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

func AddSubscriptionPlan(
	addPlan *models.AddPlansModel) (int64, error) {
	stmt := `
	INSERT INTO subscription_plans (
		name, description, price, billing_cycle, max_users,
		is_active
	)
	VALUES (?, ?, ?, ?, ?, ?)
`
	result, err := config.DB.Exec(stmt, addPlan.Name, addPlan.Description,
		addPlan.Price, addPlan.BillingCycle,
		addPlan.MaxUsers, addPlan.IsActive) // Using userID as company_id
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("could not insert New Plan: %v", err)
	}

	jobID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to add new plans ID: %v", err)
	}

	return jobID, nil
}

// GetSubscriptionPlans fetches all subscription plans
func GetSubscriptionPlans() ([]models.GetSubscriptionPlanModel, error) {
	var plans []models.GetSubscriptionPlanModel

	query := `
		SELECT plan_id, name, description, price, billing_cycle, max_users, is_active, created_at
		FROM subscription_plans
		ORDER BY created_at DESC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching subscription plans: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var plan models.GetSubscriptionPlanModel
		err := rows.Scan(
			&plan.ID,
			&plan.Name,
			&plan.Description,
			&plan.Price,
			&plan.BillingCycle,
			&plan.MaxUsers,
			&plan.IsActive,
			&plan.CreateAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning subscription plan: %v", err)
		}
		plans = append(plans, plan)
	}

	if len(plans) == 0 {
		return []models.GetSubscriptionPlanModel{}, nil
	}

	return plans, nil
}

func UpdatePaymentStatus(
	subscriptionPaymentUpdateModel *models.SubscriptionPayment) (int64, error) {
	stmt := `
	INSERT INTO subscription_payments (
		subscription_id, amount, company_id, status,
		payment_method, transaction_id
	)
	VALUES (?, ?, ?, ?, ?, ?)
`
	result, err := config.DB.Exec(stmt, 
		subscriptionPaymentUpdateModel.SubscriptionId,
		subscriptionPaymentUpdateModel.Amount,
		subscriptionPaymentUpdateModel.CompanyId,
		subscriptionPaymentUpdateModel.Status, subscriptionPaymentUpdateModel.PaymentMethod, subscriptionPaymentUpdateModel.TransactionId) // Using userID as company_id
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("could not insert New New Payment status: %v", err)
	}

	jobID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to add new Update Payment: %v", err)
	}

	return jobID, nil
}



func SubscribePlan(
	subscriptionPaymentUpdateModel *models.PlanSubscriptionModel) (int64, error) {
		fmt.Println(subscriptionPaymentUpdateModel.Status)
	stmt := `
	INSERT INTO company_subscriptions (
		company_id, plan_id, start_date, end_date,
		status, auto_renew
	)
	VALUES (?, ?, ?, ?, ?, ?)
`
	result, err := config.DB.Exec(stmt, 
		subscriptionPaymentUpdateModel.CompanyId,
		subscriptionPaymentUpdateModel.PlanId,
		subscriptionPaymentUpdateModel.StartDate,
		subscriptionPaymentUpdateModel.EndDate, 
		subscriptionPaymentUpdateModel.Status, 
		subscriptionPaymentUpdateModel.AutoRenew) // Using userID as company_id
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("could not insert into company_subscriptions DB: %v", err)
	}

	jobID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to add new Subscribe Plan: %v", err)
	}

	return jobID, nil
}


func CheckSubscriptionPlanStatus(companyID int) (string, error) {
	query := `
		SELECT status
		FROM company_subscriptions
		WHERE company_id = ?
		  AND status = 'active'
		  AND CURDATE() BETWEEN start_date AND IFNULL(end_date, CURDATE())
		ORDER BY created_at DESC
		LIMIT 1;
	`

	var status string
	err := config.DB.QueryRow(query, companyID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			// No active subscription found
			return "inactive", nil
		}
		return "inactive", fmt.Errorf("failed to check subscription status: %v", err)
	}

	return status, nil
}
