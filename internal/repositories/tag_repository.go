package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type TagRepository interface {
	FindAll() ([]models.Tag, error)
	FindByID(id uint) (*models.Tag, error)
	FindBySlug(slug string) (*models.Tag, error)
	FindByIDs(ids []uint) ([]models.Tag, error)
	Create(tag *models.Tag) error
	Update(tag *models.Tag) error
	Delete(id uint) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	if err := r.db.Order("name ASC").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) FindByID(id uint) (*models.Tag, error) {
	var tag models.Tag
	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) FindBySlug(slug string) (*models.Tag, error) {
	var tag models.Tag
	if err := r.db.Where("slug = ?", slug).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) FindByIDs(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag
	if err := r.db.Find(&tags, ids).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) Create(tag *models.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) Update(tag *models.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) Delete(id uint) error {
	return r.db.Delete(&models.Tag{}, id).Error
}
