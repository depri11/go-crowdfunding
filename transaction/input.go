package transaction

import "github.com/depri11/go-crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"status"`
	OrderID int `json:"order_id"`
	PaymentType string `json:"payment_type"`
	FraudStatus string `json:"fraud_status"`
}