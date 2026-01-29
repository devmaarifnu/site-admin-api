package middlewares

import (
	"site-admin-api/internal/utils"
	"site-admin-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// PermissionMiddleware checks if user has required permission
func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user permissions from context (set by AuthMiddleware)
		permissions := GetUserPermissions(c)

		// Check if user has the required permission
		if !utils.HasPermission(permissions, requiredPermission) {
			response.Forbidden(c, "You don't have permission to access this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RoleMiddleware checks if user has one of the required roles
func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := GetUserRole(c)

		hasRole := false
		for _, role := range requiredRoles {
			if userRole == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			response.Forbidden(c, "You don't have permission to access this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}
