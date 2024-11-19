package handlers

import (
	"net/http"
	"subscription-tracker/internal/models"
	"subscription-tracker/internal/services"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req services.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	category, err := h.categoryService.Create(&req, userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(category))
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	categories, err := h.categoryService.GetAll(userID.(models.ULID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(categories))
}

func (h *CategoryHandler) Update(c *gin.Context) {
	// Get category ID from URL
	var categoryID models.ULID
	if err := categoryID.UnmarshalJSON([]byte(`"` + c.Param("id") + `"`)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid category ID"))
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("User not found in context"))
		return
	}

	// Parse request body
	var req services.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	// Update the category
	category, err := h.categoryService.Update(categoryID, &req, userID.(models.ULID))
	if err != nil {
		switch err.Error() {
		case "category not found":
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		case "cannot edit default category":
			c.JSON(http.StatusForbidden, utils.ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(category))
}
