package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type PageRepository interface {
	FindAll(filters map[string]interface{}) ([]models.Page, error)
	FindByID(id uint) (*models.Page, error)
	FindBySlug(slug string) (*models.Page, error)
	FindActive() ([]models.Page, error)
	Create(page *models.Page) error
	Update(page *models.Page) error
	Delete(id uint) error
}

type pageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) PageRepository {
	return &pageRepository{db: db}
}

func (r *pageRepository) FindAll(filters map[string]interface{}) ([]models.Page, error) {
	var pages []models.Page
	query := r.db.Model(&models.Page{}).Preload("LastEditor")

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("updated_at DESC").Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}

func (r *pageRepository) FindByID(id uint) (*models.Page, error) {
	var page models.Page
	if err := r.db.Preload("LastEditor").First(&page, id).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *pageRepository) FindBySlug(slug string) (*models.Page, error) {
	var page models.Page
	if err := r.db.Preload("LastEditor").Where("slug = ?", slug).First(&page).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *pageRepository) FindActive() ([]models.Page, error) {
	var pages []models.Page
	if err := r.db.Where("is_active = ?", true).Preload("LastEditor").Order("updated_at DESC").Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}

func (r *pageRepository) Create(page *models.Page) error {
	return r.db.Create(page).Error
}

func (r *pageRepository) Update(page *models.Page) error {
	return r.db.Save(page).Error
}

func (r *pageRepository) Delete(id uint) error {
	return r.db.Delete(&models.Page{}, id).Error
}
