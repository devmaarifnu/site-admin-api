package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// PaginationMeta contains pagination metadata
type PaginationMeta struct {
	CurrentPage   int  `json:"current_page"`
	TotalPages    int  `json:"total_pages"`
	TotalItems    int64 `json:"total_items"`
	ItemsPerPage  int  `json:"items_per_page"`
	HasNext       bool `json:"has_next"`
	HasPrev       bool `json:"has_prev"`
}

// PaginatedResponse is the response structure with pagination
type PaginatedResponse struct {
	Success    bool           `json:"success"`
	Message    string         `json:"message,omitempty"`
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// Success sends a successful response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessWithPagination sends a successful response with pagination
func SuccessWithPagination(c *gin.Context, message string, data interface{}, pagination PaginationMeta) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// Created sends a 201 created response
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// NoContent sends a 204 no content response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// BadRequest sends a 400 bad request response
func BadRequest(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: message,
		Errors:  err,
	})
}

// Unauthorized sends a 401 unauthorized response
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: message,
	})
}

// Forbidden sends a 403 forbidden response
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, Response{
		Success: false,
		Message: message,
	})
}

// NotFound sends a 404 not found response
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: message,
	})
}

// Conflict sends a 409 conflict response
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, Response{
		Success: false,
		Message: message,
	})
}

// ValidationError sends a 422 validation error response
func ValidationError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusUnprocessableEntity, Response{
		Success: false,
		Message: "Validation error",
		Errors:  errors,
	})
}

// InternalServerError sends a 500 internal server error response
func InternalServerError(c *gin.Context, message string) {
	if message == "" {
		message = "Internal server error"
	}
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: message,
	})
}

// Error sends an error response with custom status code
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
	})
}

// ErrorWithData sends an error response with data
func ErrorWithData(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   message,
		Data:    data,
	})
}
