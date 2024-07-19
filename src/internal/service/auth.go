package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nextlua/dimo/internal/models"
)

const dimoRedirectUrlFormat = "https://auth.dev.dimo.zone/auth?client_id=%s&redirect_uri=%s&scope=openid email&response_type=code"

// Auth represents auth post method
// swagger:operation POST /auth authRequest
// ---
// summary: Authentication
// description: Endpoint to authenticate
// responses:
//
//	  200:
//		  $ref: "#/responses/authResponse"
func (s *service) Auth(ctx *gin.Context, req models.AuthInput) (int, models.AuthResponse) {
	url := s.getRedirectedUrl()
	ctx.Redirect(http.StatusTemporaryRedirect, url)
	return 404, models.AuthResponse{BaseResponse: models.BaseResponse{Message: "Not found"}}
}

func (s *service) getRedirectedUrl() string {
	return fmt.Sprintf(dimoRedirectUrlFormat, s.envVariables.Service.ClientID, s.envVariables.Service.RedirectURL)
}
