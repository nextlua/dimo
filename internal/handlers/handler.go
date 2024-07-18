package handlers

import (
	"github.com/gin-gonic/gin"
	"nextlua/dimo/internal/service"
)

type Handler struct {
	Service service.Service
}

func NewHandler(srv service.Service) *Handler {
	return &Handler{Service: srv}
}

func SetHandler(engine *gin.Engine, srv service.Service) {
	h := NewHandler(srv)

	authApi := engine.Group("/")
	{
		authApi.POST("auth", h.Auth)
	}

}
