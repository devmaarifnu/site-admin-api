package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
	"site-admin-api/internal/utils"

	"gorm.io/gorm"
)

// AuthService defines authentication service methods
type AuthService interface {
	Login(email, password string) (*models.LoginResponse, error)
	RefreshToken(refreshToken string) (*models.LoginResponse, error)
	ForgotPassword(email string) error
	ResetPassword(token, newPassword string) error
	ValidateToken(token string) (*utils.JWTClaims, error)
	BlacklistToken(token string) error
}

type authService struct {
	userRepo repositories.UserRepository
	cfg      *config.Config
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repositories.UserRepository, cfg *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *authService) Login(email, password string) (*models.LoginResponse, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Verify password
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Update last login
	s.userRepo.UpdateLastLogin(user.ID)

	// Generate tokens
	permissions := utils.GetPermissionsForRole(user.Role)
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, user.Role, permissions, s.cfg.JWT.Secret, s.cfg.JWT.ExpiresHours)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateJWT(user.ID, user.Email, user.Role, permissions, s.cfg.JWT.RefreshSecret, s.cfg.JWT.RefreshExpiresHours)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		User:         user.ToResponse(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    s.cfg.JWT.ExpiresHours * 3600,
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*models.LoginResponse, error) {
	// Validate refresh token
	claims, err := utils.ValidateJWT(refreshToken, s.cfg.JWT.RefreshSecret)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Get user
	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Generate new tokens
	permissions := utils.GetPermissionsForRole(user.Role)
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, user.Role, permissions, s.cfg.JWT.Secret, s.cfg.JWT.ExpiresHours)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateJWT(user.ID, user.Email, user.Role, permissions, s.cfg.JWT.RefreshSecret, s.cfg.JWT.RefreshExpiresHours)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		User:         user.ToResponse(),
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    s.cfg.JWT.ExpiresHours * 3600,
	}, nil
}

func (s *authService) ForgotPassword(email string) error {
	// Check if user exists
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		// Don't reveal if email exists or not for security
		return nil
	}

	// Generate reset token
	token := generateSecureToken(32)

	// Create password reset record
	reset := &models.PasswordReset{
		Email: user.Email,
		Token: token,
	}

	if err := s.userRepo.CreatePasswordReset(reset); err != nil {
		return err
	}

	// TODO: Send email with reset link
	// For now, just log it (in production, integrate with email service)
	// resetLink := s.cfg.Frontend.URL + s.cfg.Frontend.PasswordResetPath + "?token=" + token
	// sendEmail(user.Email, resetLink)

	return nil
}

func (s *authService) ResetPassword(token, newPassword string) error {
	// Find reset record
	reset, err := s.userRepo.FindPasswordReset(token)
	if err != nil {
		return errors.New("invalid or expired reset token")
	}

	// Check if token is expired (24 hours)
	if time.Since(reset.CreatedAt) > 24*time.Hour {
		s.userRepo.DeletePasswordReset(reset.Email)
		return errors.New("reset token has expired")
	}

	// Get user
	user, err := s.userRepo.FindByEmail(reset.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password
	user.Password = hashedPassword
	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	// Delete reset record
	s.userRepo.DeletePasswordReset(reset.Email)

	return nil
}

func (s *authService) ValidateToken(token string) (*utils.JWTClaims, error) {
	// Check if token is blacklisted
	isBlacklisted := s.userRepo.IsTokenBlacklisted(token)
	fmt.Printf("AUTH SERVICE: Token blacklist check - isBlacklisted: %v\n", isBlacklisted)
	if isBlacklisted {
		return nil, errors.New("token has been revoked")
	}
	return utils.ValidateJWT(token, s.cfg.JWT.Secret)
}

func (s *authService) BlacklistToken(token string) error {
	fmt.Printf("AUTH SERVICE: Attempting to blacklist token\n")
	err := s.userRepo.BlacklistToken(token)
	if err != nil {
		fmt.Printf("AUTH SERVICE: Failed to blacklist token: %v\n", err)
	} else {
		fmt.Printf("AUTH SERVICE: Token blacklisted successfully\n")
	}
	return err
}

// generateSecureToken generates a secure random token
func generateSecureToken(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
