package handlers

import (
	"site-admin-api/internal/middlewares"
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication endpoints
type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService services.AuthService, userService services.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	loginResp, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.Success(c, "Login successful", loginResp)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	loginResp, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.Success(c, "Token refreshed successfully", loginResp)
}

// GetCurrentUser returns the current authenticated user
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID := middlewares.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	user, err := h.userService.GetByID(userID)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, "User retrieved successfully", user)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *gin.Context) {
	// In JWT, logout is typically handled client-side by removing the token
	// Server-side logout would require token blacklisting (optional feature)
	response.Success(c, "Logout successful", nil)
}

// ChangePassword handles password change for authenticated user
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID := middlewares.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.userService.ChangePassword(userID, &req); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "Password changed successfully", nil)
}

// ForgotPassword handles forgot password request
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.authService.ForgotPassword(req.Email); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Password reset link has been sent to your email", nil)
}

// ResetPassword handles password reset
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req models.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.authService.ResetPassword(req.Token, req.NewPassword); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "Password reset successfully", nil)
}
