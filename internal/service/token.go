package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"nextlua/dimo/internal/models"
)

const dimoTokenFormat = "https://auth.dev.dimo.zone/token"

// Token represents auth post method
// swagger:operation POST /auth authRequest
// ---
// summary: Authentication
// description: Endpoint to authenticate
// responses:
//
//	  200:
//		  $ref: "#/responses/authResponse"
func (s *service) Token(ctx *gin.Context, req models.TokenInput) (int, models.TokenResponse) {
	tokenResponse, err := s.sendTokenRequest(req.Code)
	if err != nil {
		return http.StatusInternalServerError, models.TokenResponse{BaseResponse: models.BaseResponse{Message: err.Error()}}
	}
	return http.StatusOK, models.TokenResponse{
		BaseResponse: models.BaseResponse{Message: "Success", IsSucceed: true},
		Data:         &models.TokenResponseData{AccessToken: tokenResponse.AccessToken, ExpiresIn: tokenResponse.ExpiresIn},
	}
}

func (s *service) sendTokenRequest(code string) (*models.DimoTokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", s.envVariables.Service.ClientID)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", s.envVariables.Service.RedirectURL)

	req, err := http.NewRequest("POST", dimoTokenFormat, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var tokenResponse models.DimoTokenResponse
	if err = json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
