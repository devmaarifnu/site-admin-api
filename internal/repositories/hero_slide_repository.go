package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type HeroSlideRepository interface {
	FindAll(filters map[string]interface{}) ([]models.HeroSlide, error)
	FindByID(id uint) (*models.HeroSlide, error)
	FindActive() ([]models.HeroSlide, error)
	Create(slide *models.HeroSlide) error
	Update(slide *models.HeroSlide) error
	Delete(id uint) error
}

type heroSlideRepository struct {
	db *gorm.DB
}

func NewHeroSlideRepository(db *gorm.DB) HeroSlideRepository {
	return &heroSlideRepository{db: db}
}

func (r *heroSlideRepository) FindAll(filters map[string]interface{}) ([]models.HeroSlide, error) {
	var slides []models.HeroSlide
	query := r.db.Model(&models.HeroSlide{})

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("order_number ASC").Find(&slides).Error; err != nil {
		return nil, err
	}
	return slides, nil
}

func (r *heroSlideRepository) FindByID(id uint) (*models.HeroSlide, error) {
	var slide models.HeroSlide
	if err := r.db.First(&slide, id).Error; err != nil {
		return nil, err
	}
	return &slide, nil
}

func (r *heroSlideRepository) FindActive() ([]models.HeroSlide, error) {
	var slides []models.HeroSlide
	if err := r.db.Where("is_active = ?", true).Order("order_number ASC").Find(&slides).Error; err != nil {
		return nil, err
	}
	return slides, nil
}

func (r *heroSlideRepository) Create(slide *models.HeroSlide) error {
	return r.db.Create(slide).Error
}

func (r *heroSlideRepository) Update(slide *models.HeroSlide) error {
	return r.db.Save(slide).Error
}

func (r *heroSlideRepository) Delete(id uint) error {
	return r.db.Delete(&models.HeroSlide{}, id).Error
}
