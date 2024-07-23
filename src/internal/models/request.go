package models

type AuthInput struct {
	Key string `json:"key"`
}

type AuthRedirectInput struct {
	Code string `json:"code" example:"asd"` //code
}

type TokenInput struct {
	Code string `json:"code" example:"asd"` //code
}
