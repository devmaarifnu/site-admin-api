package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

// OrganizationPosition Repository
type OrganizationPositionRepository interface {
	FindAll(filters map[string]interface{}) ([]models.OrganizationPosition, error)
	FindByID(id uint) (*models.OrganizationPosition, error)
	FindByType(positionType string) ([]models.OrganizationPosition, error)
	Create(position *models.OrganizationPosition) error
	Update(position *models.OrganizationPosition) error
	Delete(id uint) error
}

type organizationPositionRepository struct {
	db *gorm.DB
}

func NewOrganizationPositionRepository(db *gorm.DB) OrganizationPositionRepository {
	return &organizationPositionRepository{db: db}
}

func (r *organizationPositionRepository) FindAll(filters map[string]interface{}) ([]models.OrganizationPosition, error) {
	var positions []models.OrganizationPosition
	query := r.db.Model(&models.OrganizationPosition{}).Preload("Parent")

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC").Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}

func (r *organizationPositionRepository) FindByID(id uint) (*models.OrganizationPosition, error) {
	var position models.OrganizationPosition
	if err := r.db.Preload("Parent").First(&position, id).Error; err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *organizationPositionRepository) FindByType(positionType string) ([]models.OrganizationPosition, error) {
	var positions []models.OrganizationPosition
	if err := r.db.Where("position_type = ? AND is_active = ?", positionType, true).
		Order("order_number ASC").Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}

func (r *organizationPositionRepository) Create(position *models.OrganizationPosition) error {
	return r.db.Create(position).Error
}

func (r *organizationPositionRepository) Update(position *models.OrganizationPosition) error {
	return r.db.Save(position).Error
}

func (r *organizationPositionRepository) Delete(id uint) error {
	return r.db.Delete(&models.OrganizationPosition{}, id).Error
}

// BoardMember Repository
type BoardMemberRepository interface {
	FindAll(filters map[string]interface{}) ([]models.BoardMember, error)
	FindByID(id uint) (*models.BoardMember, error)
	FindByPeriod(start, end int) ([]models.BoardMember, error)
	FindActive() ([]models.BoardMember, error)
	Create(member *models.BoardMember) error
	Update(member *models.BoardMember) error
	Delete(id uint) error
}

type boardMemberRepository struct {
	db *gorm.DB
}

func NewBoardMemberRepository(db *gorm.DB) BoardMemberRepository {
	return &boardMemberRepository{db: db}
}

func (r *boardMemberRepository) FindAll(filters map[string]interface{}) ([]models.BoardMember, error) {
	var members []models.BoardMember
	query := r.db.Model(&models.BoardMember{}).Preload("Position")

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

func (r *boardMemberRepository) FindByID(id uint) (*models.BoardMember, error) {
	var member models.BoardMember
	if err := r.db.Preload("Position").First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *boardMemberRepository) FindByPeriod(start, end int) ([]models.BoardMember, error) {
	var members []models.BoardMember
	if err := r.db.Where("period_start = ? AND period_end = ?", start, end).
		Preload("Position").Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *boardMemberRepository) FindActive() ([]models.BoardMember, error) {
	var members []models.BoardMember
	if err := r.db.Where("is_active = ?", true).Preload("Position").Order("order_number ASC").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *boardMemberRepository) Create(member *models.BoardMember) error {
	return r.db.Create(member).Error
}

func (r *boardMemberRepository) Update(member *models.BoardMember) error {
	return r.db.Save(member).Error
}

func (r *boardMemberRepository) Delete(id uint) error {
	return r.db.Delete(&models.BoardMember{}, id).Error
}

// Pengurus Repository
type PengurusRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Pengurus, error)
	FindByID(id uint) (*models.Pengurus, error)
	FindByCategory(category string) ([]models.Pengurus, error)
	FindByPeriod(start, end int) ([]models.Pengurus, error)
	FindActive() ([]models.Pengurus, error)
	Create(pengurus *models.Pengurus) error
	Update(pengurus *models.Pengurus) error
	Delete(id uint) error
}

type pengurusRepository struct {
	db *gorm.DB
}

func NewPengurusRepository(db *gorm.DB) PengurusRepository {
	return &pengurusRepository{db: db}
}

func (r *pengurusRepository) FindAll(filters map[string]interface{}) ([]models.Pengurus, error) {
	var pengurusList []models.Pengurus
	query := r.db.Model(&models.Pengurus{})

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC").Find(&pengurusList).Error; err != nil {
		return nil, err
	}
	return pengurusList, nil
}

func (r *pengurusRepository) FindByID(id uint) (*models.Pengurus, error) {
	var pengurus models.Pengurus
	if err := r.db.First(&pengurus, id).Error; err != nil {
		return nil, err
	}
	return &pengurus, nil
}

func (r *pengurusRepository) FindByCategory(category string) ([]models.Pengurus, error) {
	var pengurusList []models.Pengurus
	if err := r.db.Where("kategori = ? AND is_active = ?", category, true).
		Order("order_number ASC").Find(&pengurusList).Error; err != nil {
		return nil, err
	}
	return pengurusList, nil
}

func (r *pengurusRepository) FindByPeriod(start, end int) ([]models.Pengurus, error) {
	var pengurusList []models.Pengurus
	if err := r.db.Where("periode_mulai = ? AND periode_selesai = ?", start, end).
		Order("order_number ASC").Find(&pengurusList).Error; err != nil {
		return nil, err
	}
	return pengurusList, nil
}

func (r *pengurusRepository) FindActive() ([]models.Pengurus, error) {
	var pengurusList []models.Pengurus
	if err := r.db.Where("is_active = ?", true).Order("order_number ASC").Find(&pengurusList).Error; err != nil {
		return nil, err
	}
	return pengurusList, nil
}

func (r *pengurusRepository) Create(pengurus *models.Pengurus) error {
	return r.db.Create(pengurus).Error
}

func (r *pengurusRepository) Update(pengurus *models.Pengurus) error {
	return r.db.Save(pengurus).Error
}

func (r *pengurusRepository) Delete(id uint) error {
	return r.db.Delete(&models.Pengurus{}, id).Error
}

// Department Repository
type DepartmentRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Department, error)
	FindByID(id uint) (*models.Department, error)
	FindActive() ([]models.Department, error)
	Create(dept *models.Department) error
	Update(dept *models.Department) error
	Delete(id uint) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) FindAll(filters map[string]interface{}) ([]models.Department, error) {
	var departments []models.Department
	query := r.db.Model(&models.Department{})

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC").Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *departmentRepository) FindByID(id uint) (*models.Department, error) {
	var dept models.Department
	if err := r.db.First(&dept, id).Error; err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *departmentRepository) FindActive() ([]models.Department, error) {
	var departments []models.Department
	if err := r.db.Where("is_active = ?", true).Order("order_number ASC").Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *departmentRepository) Create(dept *models.Department) error {
	return r.db.Create(dept).Error
}

func (r *departmentRepository) Update(dept *models.Department) error {
	return r.db.Save(dept).Error
}

func (r *departmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}
