package handlers

import (
	"net/http"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/services"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	subscriptionService *services.SubscriptionService
}

func NewSubscriptionHandler(subscriptionService *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

func (h *SubscriptionHandler) Create(c *gin.Context) {
	var req services.CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	subscription, err := h.subscriptionService.Create(&req, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(subscription))
}
