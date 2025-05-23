package models


type GetSubscriptionPlanModel struct {
	ID			 string    `json:"plan_id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	Price        float64   `json:"price" db:"price"`
	BillingCycle string    `json:"billing_cycle" db:"billing_cycle"` // "monthly" or "yearly"
	MaxUsers     *int      `json:"max_users,omitempty" db:"max_users"` // nullable
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreateAt     string    `json:"created_at"`
}

