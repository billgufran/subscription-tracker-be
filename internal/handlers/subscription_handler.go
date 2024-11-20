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

func (h *SubscriptionHandler) GetAll(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	subscriptions, err := h.subscriptionService.GetAll(userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(subscriptions))
}

func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	var subscriptionID models.ULID
	if err := subscriptionID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subscription ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	subscription, err := h.subscriptionService.GetByID(subscriptionID, userID.(models.ULID))
	if err != nil {
		if err.Error() == "subscription not found" {
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(subscription))
}

func (h *SubscriptionHandler) GetByCategory(c *gin.Context) {
	var categoryID models.ULID
	if err := categoryID.UnmarshalJSON([]byte(`"` + c.Param("categoryId") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid category ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	subscriptions, err := h.subscriptionService.GetByCategory(categoryID, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(subscriptions))
}

func (h *SubscriptionHandler) GetByBillingCycle(c *gin.Context) {
	var billingCycleID models.ULID
	if err := billingCycleID.UnmarshalJSON([]byte(`"` + c.Param("billingCycleId") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid billing cycle ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	subscriptions, err := h.subscriptionService.GetByBillingCycle(billingCycleID, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(subscriptions))
}

func (h *SubscriptionHandler) GetByPaymentMethod(c *gin.Context) {
	var paymentMethodID models.ULID
	if err := paymentMethodID.UnmarshalJSON([]byte(`"` + c.Param("paymentMethodId") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid payment method ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	subscriptions, err := h.subscriptionService.GetByPaymentMethod(paymentMethodID, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(subscriptions))
}

func (h *SubscriptionHandler) Update(c *gin.Context) {
	var subscriptionID models.ULID
	if err := subscriptionID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subscription ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	var req services.UpdateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	subscription, err := h.subscriptionService.Update(subscriptionID, &req, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "subscription not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(subscription))
}

func (h *SubscriptionHandler) Delete(c *gin.Context) {
	var subscriptionID models.ULID
	if err := subscriptionID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subscription ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	err := h.subscriptionService.Delete(subscriptionID, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "subscription not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}
