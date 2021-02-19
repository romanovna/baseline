package docs

import (
	"crud-api/models"
)

// Package classification posts crud_api.
//
// Documentation of our crud_api API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
//
// swagger:meta

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Returns list of entities
// swagger:response postListResponse
type postListResponse struct {
	// in: body
	Body []models.Post
}

// Returns single entity
// swagger:response postSingleResponse
type postSingleResponse struct {
	// in: body
	Body models.Post
}

// swagger:parameters updatePost
type postIDParamsWrapper struct {
	// The id of the post for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updatePost createPost
type productParamsWrapper struct {
	// Post data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// The body of the post for which the operation relates
	// in: body
	// required: true
	Body models.Post
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}
