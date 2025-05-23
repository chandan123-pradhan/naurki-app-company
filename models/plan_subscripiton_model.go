
package models
type PlanSubscriptionModel struct {
	PlanId int `json:"plan_id"`
	StartDate string `json:"start_date"`
	EndDate  string `json:"end_date"`
	Status string `json:"status"`
	AutoRenew bool `json:"auto_renew"`
	CompanyId int `json:"company_id"`
}