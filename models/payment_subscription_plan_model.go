package models

type SubscriptionPayment struct {
	SubscriptionId string     `json:"subscription_id"`
	Amount         float64 `json:"amount"`
	CompanyId      int     `json:"company_id"`
	Status         string  `json:"status"`
	PaymentMethod  string  `json:"payment_method"`
	TransactionId  string  `json:"transaction_id"`
}
