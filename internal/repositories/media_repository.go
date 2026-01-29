package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type MediaRepository interface {
	FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.Media, int64, error)
	FindByID(id uint) (*models.Media, error)
	FindByFolder(folder string, page, limit int) ([]models.Media, int64, error)
	FindByType(fileType string, page, limit int) ([]models.Media, int64, error)
	Create(media *models.Media) error
	Update(media *models.Media) error
	Delete(id uint) error
}

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

func (r *mediaRepository) FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.Media, int64, error) {
	var mediaList []models.Media
	var total int64

	query := r.db.Model(&models.Media{}).Preload("Uploader")

	if search != "" {
		query = query.Where("file_name LIKE ? OR original_name LIKE ? OR alt_text LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&mediaList).Error; err != nil {
		return nil, 0, err
	}

	return mediaList, total, nil
}

func (r *mediaRepository) FindByID(id uint) (*models.Media, error) {
	var media models.Media
	if err := r.db.Preload("Uploader").First(&media, id).Error; err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepository) FindByFolder(folder string, page, limit int) ([]models.Media, int64, error) {
	var mediaList []models.Media
	var total int64

	query := r.db.Model(&models.Media{}).Where("folder = ?", folder).Preload("Uploader")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&mediaList).Error; err != nil {
		return nil, 0, err
	}

	return mediaList, total, nil
}

func (r *mediaRepository) FindByType(fileType string, page, limit int) ([]models.Media, int64, error) {
	var mediaList []models.Media
	var total int64

	query := r.db.Model(&models.Media{}).Where("file_type = ?", fileType).Preload("Uploader")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&mediaList).Error; err != nil {
		return nil, 0, err
	}

	return mediaList, total, nil
}

func (r *mediaRepository) Create(media *models.Media) error {
	return r.db.Create(media).Error
}

func (r *mediaRepository) Update(media *models.Media) error {
	return r.db.Save(media).Error
}

func (r *mediaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Media{}, id).Error
}
