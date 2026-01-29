package handlers

import (
	"strconv"

	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/internal/utils"
	"site-admin-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user management endpoints
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAll returns all users with pagination
func (h *UserHandler) GetAll(c *gin.Context) {
	// Get pagination params
	params := utils.GetPaginationParams(c, 20, 100)

	// Get search and filters
	search := utils.GetSearchQuery(c)
	filters := utils.GetFilterParams(c)

	// Get users
	users, total, err := h.userService.GetAll(params.Page, params.Limit, search, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	// Calculate pagination meta
	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)

	response.SuccessWithPagination(c, "Users retrieved successfully", users, pagination)
}

// GetByID returns a user by ID
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	user, err := h.userService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, "User retrieved successfully", user)
}

// Create creates a new user
func (h *UserHandler) Create(c *gin.Context) {
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	user, err := h.userService.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "User created successfully", user)
}

// Update updates a user
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	var req models.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	user, err := h.userService.Update(uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "User updated successfully", user)
}

// Delete deletes a user
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	if err := h.userService.Delete(uint(id)); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "User deleted successfully", nil)
}

// UpdateStatus updates user status
func (h *UserHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	var req struct {
		Status string `json:"status" binding:"required,oneof=active inactive"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.userService.UpdateStatus(uint(id), req.Status); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "User status updated successfully", nil)
}
