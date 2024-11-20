package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHttpError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		switch appErr.Code {
		case CodeNotFound:
			c.JSON(http.StatusNotFound, ErrorResponse(err.Error()))
		case CodeValidation:
			c.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		case CodeForbidden:
			c.JSON(http.StatusForbidden, ErrorResponse(err.Error()))
		case CodeUnauthorized:
			c.JSON(http.StatusUnauthorized, ErrorResponse(err.Error()))
		case CodeDuplicate:
			c.JSON(http.StatusConflict, ErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse(err.Error()))
		}
		return
	}
	// Handle non-AppError errors
	c.JSON(http.StatusInternalServerError, ErrorResponse(err.Error()))
}
