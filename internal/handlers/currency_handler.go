package handlers

import (
	"net/http"
	"subscription-tracker/internal/services"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	currencyService *services.CurrencyService
}

func NewCurrencyHandler(currencyService *services.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: currencyService,
	}
}

func (h *CurrencyHandler) GetAll(c *gin.Context) {
	currencies, err := h.currencyService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(currencies))
}
