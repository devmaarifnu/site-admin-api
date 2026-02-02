package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the users table
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	Role            string         `gorm:"type:enum('super_admin','admin','editor','viewer');default:'viewer'" json:"role"`
	Avatar          *string        `gorm:"type:varchar(500)" json:"avatar"`
	Phone           *string        `gorm:"type:varchar(20)" json:"phone"`
	IsActive        bool           `gorm:"default:true" json:"is_active"`
	LastLoginAt     *time.Time     `json:"last_login_at"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	RememberToken   *string        `gorm:"type:varchar(100)" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name
func (User) TableName() string {
	return "users"
}

// UserCreateRequest represents the request to create a user
type UserCreateRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8"`
	Role     string  `json:"role" binding:"required,oneof=super_admin admin editor viewer"`
	Phone    *string `json:"phone"`
	Avatar   *string `json:"avatar"`
}

// UserUpdateRequest represents the request to update a user
type UserUpdateRequest struct {
	Name   string  `json:"name"`
	Email  string  `json:"email" binding:"omitempty,email"`
	Role   string  `json:"role" binding:"omitempty,oneof=super_admin admin editor viewer"`
	Phone  *string `json:"phone"`
	Avatar *string `json:"avatar"`
}

// UserResponse represents the user response (without sensitive data)
type UserResponse struct {
	ID              uint       `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Role            string     `json:"role"`
	Avatar          *string    `json:"avatar"`
	Phone           *string    `json:"phone"`
	IsActive        bool       `json:"is_active"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		Role:            u.Role,
		Avatar:          u.Avatar,
		Phone:           u.Phone,
		IsActive:        u.IsActive,
		LastLoginAt:     u.LastLoginAt,
		EmailVerifiedAt: u.EmailVerifiedAt,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}

// LoginRequest represents the login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int          `json:"expires_in"`
}

// RefreshTokenRequest represents the refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// ChangePasswordRequest represents the change password request
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// ForgotPasswordRequest represents the forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ResetPasswordRequest represents the reset password request
type ResetPasswordRequest struct{
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// PasswordReset represents the password_resets table
type PasswordReset struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"size:255;not null;index" json:"email"`
	Token     string    `gorm:"size:255;not null;index" json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name
func (PasswordReset) TableName() string {
	return "password_resets"
}

// PersonalAccessToken represents the personal_access_tokens table (used for token blacklist)
type PersonalAccessToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"type:text;not null;index" json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name
func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens"
}
