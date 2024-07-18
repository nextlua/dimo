package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nextlua/dimo/internal/models"
)

// AuthRedirect represents auth redirect post method
// swagger:operation POST /auth_redirect authRedirectRequest
// ---
// summary: Authentication Redirect
// description: Endpoint to authenticate
// responses:
//
//	  200:
//		  $ref: "#/responses/authResponse"
func (s *service) AuthRedirect(ctx *gin.Context, req models.AuthRedirectInput) (int, models.AuthRedirectResponse) {
	code := req.Code
	return http.StatusOK, models.AuthRedirectResponse{
		BaseResponse: models.BaseResponse{
			Message: "Success",
		},
		Data: &models.AuthRedirectResponseData{
			Code: code,
		},
	}
}
