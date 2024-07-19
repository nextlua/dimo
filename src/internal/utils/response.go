package utils

import "nextlua/dimo/internal/models"

func GenerateErrorResponse(code int, message string) (int, any) {
	return code, &models.ErrorResponse{
		BaseResponse: models.BaseResponse{
			IsSucceed: false,
			Message:   message,
		},
		Data: nil,
	}
}
