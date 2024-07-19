package models

type AuthInput struct{}

type AuthRedirectInput struct {
	Code string `json:"code" example:"asd"` //code
}

type TokenInput struct {
	Code string `json:"code" example:"asd"` //code
}
