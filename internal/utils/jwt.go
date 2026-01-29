package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents the JWT claims structure
type JWTClaims struct {
	UserID      uint     `json:"user_id"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID uint, email, role string, permissions []string, secret string, expiresHours int) (string, error) {
	claims := JWTClaims{
		UserID:      userID,
		Email:       email,
		Role:        role,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateJWT validates a JWT token and returns the claims
func ValidateJWT(tokenString string, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetPermissionsForRole returns the permissions for a given role
func GetPermissionsForRole(role string) []string {
	permissionsMap := map[string][]string{
		"super_admin": {
			"users.view", "users.create", "users.update", "users.delete",
			"news.view", "news.create", "news.update", "news.delete",
			"opinions.view", "opinions.create", "opinions.update", "opinions.delete",
			"documents.view", "documents.create", "documents.update", "documents.delete",
			"hero_slides.view", "hero_slides.create", "hero_slides.update", "hero_slides.delete",
			"organization.view", "organization.create", "organization.update", "organization.delete",
			"pages.view", "pages.update",
			"events.view", "events.create", "events.update", "events.delete",
			"media.view", "media.upload", "media.update", "media.delete",
			"categories.view", "categories.create", "categories.update", "categories.delete",
			"tags.view", "tags.create", "tags.update", "tags.delete",
			"contact_messages.view", "contact_messages.update", "contact_messages.delete",
			"settings.view", "settings.update",
			"analytics.view",
			"activity_logs.view", "activity_logs.delete",
		},
		"admin": {
			"news.view", "news.create", "news.update", "news.delete",
			"opinions.view", "opinions.create", "opinions.update", "opinions.delete",
			"documents.view", "documents.create", "documents.update", "documents.delete",
			"hero_slides.view", "hero_slides.create", "hero_slides.update", "hero_slides.delete",
			"organization.view", "organization.create", "organization.update", "organization.delete",
			"pages.view", "pages.update",
			"events.view", "events.create", "events.update", "events.delete",
			"media.view", "media.upload", "media.update", "media.delete",
			"categories.view", "categories.create", "categories.update", "categories.delete",
			"tags.view", "tags.create", "tags.update", "tags.delete",
			"contact_messages.view", "contact_messages.update", "contact_messages.delete",
			"settings.view", "settings.update",
			"analytics.view",
		},
		"redaktur": {
			"news.view", "news.create", "news.update", "news.delete",
			"opinions.view", "opinions.create", "opinions.update", "opinions.delete",
			"media.view", "media.upload",
		},
	}

	if permissions, ok := permissionsMap[role]; ok {
		return permissions
	}
	return []string{}
}

// HasPermission checks if a user has a specific permission
func HasPermission(userPermissions []string, requiredPermission string) bool {
	for _, permission := range userPermissions {
		if permission == requiredPermission {
			return true
		}
	}
	return false
}
