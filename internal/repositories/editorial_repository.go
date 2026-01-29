package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

// EditorialTeam Repository
type EditorialTeamRepository interface {
	FindAll(filters map[string]interface{}) ([]models.EditorialTeam, error)
	FindByID(id uint) (*models.EditorialTeam, error)
	FindByRoleType(roleType string) ([]models.EditorialTeam, error)
	FindActive() ([]models.EditorialTeam, error)
	Create(member *models.EditorialTeam) error
	Update(member *models.EditorialTeam) error
	Delete(id uint) error
}

type editorialTeamRepository struct {
	db *gorm.DB
}

func NewEditorialTeamRepository(db *gorm.DB) EditorialTeamRepository {
	return &editorialTeamRepository{db: db}
}

func (r *editorialTeamRepository) FindAll(filters map[string]interface{}) ([]models.EditorialTeam, error) {
	var members []models.EditorialTeam
	query := r.db.Model(&models.EditorialTeam{})

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *editorialTeamRepository) FindByID(id uint) (*models.EditorialTeam, error) {
	var member models.EditorialTeam
	if err := r.db.First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *editorialTeamRepository) FindByRoleType(roleType string) ([]models.EditorialTeam, error) {
	var members []models.EditorialTeam
	if err := r.db.Where("role_type = ? AND is_active = ?", roleType, true).
		Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *editorialTeamRepository) FindActive() ([]models.EditorialTeam, error) {
	var members []models.EditorialTeam
	if err := r.db.Where("is_active = ?", true).Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *editorialTeamRepository) Create(member *models.EditorialTeam) error {
	return r.db.Create(member).Error
}

func (r *editorialTeamRepository) Update(member *models.EditorialTeam) error {
	return r.db.Save(member).Error
}

func (r *editorialTeamRepository) Delete(id uint) error {
	return r.db.Delete(&models.EditorialTeam{}, id).Error
}

// EditorialCouncil Repository
type EditorialCouncilRepository interface {
	FindAll(filters map[string]interface{}) ([]models.EditorialCouncil, error)
	FindByID(id uint) (*models.EditorialCouncil, error)
	FindActive() ([]models.EditorialCouncil, error)
	Create(member *models.EditorialCouncil) error
	Update(member *models.EditorialCouncil) error
	Delete(id uint) error
}

type editorialCouncilRepository struct {
	db *gorm.DB
}

func NewEditorialCouncilRepository(db *gorm.DB) EditorialCouncilRepository {
	return &editorialCouncilRepository{db: db}
}

func (r *editorialCouncilRepository) FindAll(filters map[string]interface{}) ([]models.EditorialCouncil, error) {
	var members []models.EditorialCouncil
	query := r.db.Model(&models.EditorialCouncil{})

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *editorialCouncilRepository) FindByID(id uint) (*models.EditorialCouncil, error) {
	var member models.EditorialCouncil
	if err := r.db.First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *editorialCouncilRepository) FindActive() ([]models.EditorialCouncil, error) {
	var members []models.EditorialCouncil
	if err := r.db.Where("is_active = ?", true).Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *editorialCouncilRepository) Create(member *models.EditorialCouncil) error {
	return r.db.Create(member).Error
}

func (r *editorialCouncilRepository) Update(member *models.EditorialCouncil) error {
	return r.db.Save(member).Error
}

func (r *editorialCouncilRepository) Delete(id uint) error {
	return r.db.Delete(&models.EditorialCouncil{}, id).Error
}
