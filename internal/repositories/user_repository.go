package repositories

import (
	"site-admin-api/internal/models"
	"time"

	"gorm.io/gorm"
)

// UserRepository defines methods for user data access
type UserRepository interface {
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.User, int64, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
	UpdateLastLogin(id uint) error
	CreatePasswordReset(reset *models.PasswordReset) error
	FindPasswordReset(token string) (*models.PasswordReset, error)
	DeletePasswordReset(email string) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{})

	// Apply search
	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Apply filters
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) UpdateLastLogin(id uint) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("last_login_at", now).Error
}

func (r *userRepository) CreatePasswordReset(reset *models.PasswordReset) error {
	// Delete existing resets for this email
	r.db.Where("email = ?", reset.Email).Delete(&models.PasswordReset{})
	return r.db.Create(reset).Error
}

func (r *userRepository) FindPasswordReset(token string) (*models.PasswordReset, error) {
	var reset models.PasswordReset
	if err := r.db.Where("token = ?", token).First(&reset).Error; err != nil {
		return nil, err
	}
	return &reset, nil
}

func (r *userRepository) DeletePasswordReset(email string) error {
	return r.db.Where("email = ?", email).Delete(&models.PasswordReset{}).Error
}
