package handlers

import (
	"net/http"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/services"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

type PaymentMethodHandler struct {
	paymentMethodService *services.PaymentMethodService
}

func NewPaymentMethodHandler(paymentMethodService *services.PaymentMethodService) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		paymentMethodService: paymentMethodService,
	}
}

func (h *PaymentMethodHandler) Create(c *gin.Context) {
	var req services.CreatePaymentMethodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	paymentMethod, err := h.paymentMethodService.Create(&req, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(paymentMethod))
}

func (h *PaymentMethodHandler) GetAll(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	paymentMethods, err := h.paymentMethodService.GetAll(userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(paymentMethods))
}

func (h *PaymentMethodHandler) Update(c *gin.Context) {
	var paymentMethodID models.ULID
	if err := paymentMethodID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid payment method ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	var req services.UpdatePaymentMethodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	paymentMethod, err := h.paymentMethodService.Update(paymentMethodID, &req, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "payment method not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(paymentMethod))
}

func (h *PaymentMethodHandler) Delete(c *gin.Context) {
	var paymentMethodID models.ULID
	if err := paymentMethodID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid payment method ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	err := h.paymentMethodService.Delete(paymentMethodID, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "payment method not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}
