package services

import (
	"errors"

	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
	"site-admin-api/internal/utils"

	"gorm.io/gorm"
)

// UserService defines user service methods
type UserService interface {
	GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.UserResponse, int64, error)
	GetByID(id uint) (*models.UserResponse, error)
	Create(req *models.UserCreateRequest) (*models.UserResponse, error)
	Update(id uint, req *models.UserUpdateRequest) (*models.UserResponse, error)
	Delete(id uint) error
	UpdateStatus(id uint, status string) error
	ChangePassword(id uint, req *models.ChangePasswordRequest) error
}

type userService struct {
	userRepo repositories.UserRepository
	cfg      *config.Config
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository, cfg *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *userService) GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.UserResponse, int64, error) {
	users, total, err := s.userRepo.FindAll(page, limit, search, filters)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	return responses, total, nil
}

func (s *userService) GetByID(id uint) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) Create(req *models.UserCreateRequest) (*models.UserResponse, error) {
	// Check if email already exists
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
		Avatar:   req.Avatar,
		Phone:    req.Phone,
		IsActive: true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) Update(id uint, req *models.UserUpdateRequest) (*models.UserResponse, error) {
	// Get existing user
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if email is being changed and already exists
	if req.Email != "" && req.Email != user.Email {
		_, err := s.userRepo.FindByEmail(req.Email)
		if err == nil {
			return nil, errors.New("email already exists")
		}
	}

	// Update fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.Avatar != nil {
		user.Avatar = req.Avatar
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) Delete(id uint) error {
	// Check if user exists
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.userRepo.Delete(id)
}

func (s *userService) UpdateStatus(id uint, status string) error {
	// Validate status
	if status != "active" && status != "inactive" {
		return errors.New("invalid status")
	}

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	user.IsActive = (status == "active")
	return s.userRepo.Update(user)
}

func (s *userService) ChangePassword(id uint, req *models.ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Verify old password
	if !utils.CheckPassword(req.OldPassword, user.Password) {
		return errors.New("old password is incorrect")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.Update(user)
}
