package handlers

import (
	"github.com/gin-gonic/gin"
	"nextlua/dimo/internal/models"
	"nextlua/dimo/internal/utils"
)

func (h *Handler) Auth(c *gin.Context) {
	var input models.AuthInput
	if code, res := utils.ParseJSONBody(c, &input); res != nil {
		c.JSON(code, res)
		return
	}
	c.JSON(h.Service.Auth(c, input))
}
