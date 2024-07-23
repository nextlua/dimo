package handlers

import (
	"github.com/gin-gonic/gin"
	"nextlua/dimo/internal/models"
	"nextlua/dimo/internal/utils"
)

func (h *Handler) Token(c *gin.Context) {
	var input models.TokenInput
	if code, res := utils.ParseJSONBody(c, &input); res != nil {
		c.JSON(code, res)
		return
	}

	c.JSON(h.Service.Token(c, input))
}
