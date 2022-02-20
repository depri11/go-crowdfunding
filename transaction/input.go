package transaction

import "github.com/depri11/go-crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}