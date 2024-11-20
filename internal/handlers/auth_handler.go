package handlers

import (
	"net/http"

	"subscription-tracker/internal/services"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleHttpError(c, utils.NewValidationError("body", "invalid request body"))
		return
	}

	response, err := h.authService.Login(&req)
	if err != nil {
		utils.HandleHttpError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req services.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleHttpError(c, utils.NewValidationError("body", "invalid request body"))
		return
	}

	response, err := h.authService.Register(&req)
	if err != nil {
		utils.HandleHttpError(c, err)
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(response))
}
