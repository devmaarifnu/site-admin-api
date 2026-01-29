package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
	FindBySlug(slug string) (*models.Category, error)
	FindByType(categoryType string) ([]models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll(filters map[string]interface{}) ([]models.Category, error) {
	var categories []models.Category

	query := r.db.Model(&models.Category{})

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC, name ASC").Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindBySlug(slug string) (*models.Category, error) {
	var category models.Category
	if err := r.db.Where("slug = ?", slug).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindByType(categoryType string) ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Where("type = ? AND is_active = ?", categoryType, true).
		Order("order_number ASC, name ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
