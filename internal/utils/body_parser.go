package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextlua/dimo/internal/validator"
)

func ParseJSONBody(c *gin.Context, request interface{}) (int, any) {
	if err := c.ShouldBindJSON(&request); err != nil {
		if validationErrors := validator.GetValidationErrors(err); validationErrors != "" {
			return GenerateErrorResponse(http.StatusBadRequest, validationErrors)
		} else {
			return GenerateErrorResponse(http.StatusBadRequest, err.Error())
		}
	}
	return 0, nil
}
