package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextlua/dimo/internal/models"
	"nextlua/dimo/internal/utils"
)

func (h *Handler) AuthRedirect(c *gin.Context) {
	var input models.AuthRedirectInput
	code := c.Query("code")
	name := c.Param("code")
	println(code, name)
	input.Code = code
	println(c.Request.URL.String())
	if code == "" {
		c.JSON(utils.GenerateErrorResponse(http.StatusInternalServerError, "code not exist"))
	}
	c.JSON(h.Service.AuthRedirect(c, input))
}
