package handlers

import (
	"net/http"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/services"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

type BillingCycleHandler struct {
	billingCycleService *services.BillingCycleService
}

func NewBillingCycleHandler(billingCycleService *services.BillingCycleService) *BillingCycleHandler {
	return &BillingCycleHandler{
		billingCycleService: billingCycleService,
	}
}

func (h *BillingCycleHandler) Create(c *gin.Context) {
	var req services.CreateBillingCycleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	billingCycle, err := h.billingCycleService.Create(&req, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(billingCycle))
}

func (h *BillingCycleHandler) GetAll(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	billingCycles, err := h.billingCycleService.GetAll(userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(billingCycles))
}

func (h *BillingCycleHandler) Update(c *gin.Context) {
	var billingCycleID models.ULID
	if err := billingCycleID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid billing cycle ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	var req services.UpdateBillingCycleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	billingCycle, err := h.billingCycleService.Update(billingCycleID, &req, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "billing cycle not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		case "cannot edit system-defined billing cycle":
			c.JSON(http.StatusForbidden, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(billingCycle))
}

func (h *BillingCycleHandler) Delete(c *gin.Context) {
	var billingCycleID models.ULID
	if err := billingCycleID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid billing cycle ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	err := h.billingCycleService.Delete(billingCycleID, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "billing cycle not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		case "cannot delete system-defined billing cycle":
			c.JSON(http.StatusForbidden, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}
