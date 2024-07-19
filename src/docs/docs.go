// Package docs Nextlua DIMO Authentication Service API.
//
// Documentation for Nextlua DIMO Authentication Service API
//
//	Schemes: https, http
//	BasePath: ./
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//	Security:
//	- api_key:
//
//	SecurityDefinitions:
//	api_key:
//	     type: apiKey
//	     name: KEY
//	     in: header
//
// swagger:meta
package docs

import "nextlua/dimo/internal/models"

// swagger:parameters authRequest
type reservationCancelRequest struct {
	// in: body
	Body struct {
		models.AuthInput
	}
}

// swagger:response authResponse
type reservationCancelResponse struct {
	// in:body
	ReservationCancelResponse models.AuthResponse
}
