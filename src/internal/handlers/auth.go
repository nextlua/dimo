package handlers

import (
	"github.com/gin-gonic/gin"
	"nextlua/dimo/internal/models"
)

func (h *Handler) Auth(c *gin.Context) {
	var input models.AuthInput
	c.JSON(h.Service.Auth(c, input))
}
