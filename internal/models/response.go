package models

type BaseResponse struct {
	IsSucceed bool   `json:"isSucceed"`
	Message   string `json:"message"`
}

type ErrorResponse struct {
	BaseResponse
	Data any `json:"data"`
}

type AuthResponseData struct {
	Deneme string
}

type AuthResponse struct {
	BaseResponse
	Data *AuthResponseData `json:"data"`
}

type AuthRedirectResponseData struct {
	Code string `json:"code"`
}

type AuthRedirectResponse struct {
	BaseResponse
	Data *AuthRedirectResponseData `json:"data"`
}

type TokenResponseData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type TokenResponse struct {
	BaseResponse
	Data *TokenResponseData `json:"data"`
}
