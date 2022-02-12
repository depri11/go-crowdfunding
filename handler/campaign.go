package handler

import (
	"net/http"
	"strconv"

	"github.com/depri11/go-crowdfunding/campaign"
	"github.com/depri11/go-crowdfunding/helper"
	"github.com/gin-gonic/gin"
)

//tangkap parameter
//handler ke service
//service yang menentukan apakah repository mana yang di-call
//repository : FindAll, FindByUserID
//db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaigns

func (h *campaignHandler) GetCampaigns (c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get Campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}