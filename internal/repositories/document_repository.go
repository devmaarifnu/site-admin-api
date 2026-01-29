package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type DocumentRepository interface {
	FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.Document, int64, error)
	FindByID(id uint) (*models.Document, error)
	Create(doc *models.Document) error
	Update(doc *models.Document) error
	Delete(id uint) error
	IncrementDownloadCount(id uint) error
	FindByCategory(categoryID uint, page, limit int) ([]models.Document, int64, error)
	FindPublic(page, limit int) ([]models.Document, int64, error)
}

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{db: db}
}

func (r *documentRepository) FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.Document, int64, error) {
	var documents []models.Document
	var total int64

	query := r.db.Model(&models.Document{}).Preload("Category").Preload("Uploader")

	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ? OR file_name LIKE ?",
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
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&documents).Error; err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}

func (r *documentRepository) FindByID(id uint) (*models.Document, error) {
	var doc models.Document
	if err := r.db.Preload("Category").Preload("Uploader").First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *documentRepository) Create(doc *models.Document) error {
	return r.db.Create(doc).Error
}

func (r *documentRepository) Update(doc *models.Document) error {
	return r.db.Save(doc).Error
}

func (r *documentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Document{}, id).Error
}

func (r *documentRepository) IncrementDownloadCount(id uint) error {
	return r.db.Model(&models.Document{}).Where("id = ?", id).UpdateColumn("download_count", gorm.Expr("download_count + ?", 1)).Error
}

func (r *documentRepository) FindByCategory(categoryID uint, page, limit int) ([]models.Document, int64, error) {
	var documents []models.Document
	var total int64

	query := r.db.Model(&models.Document{}).Where("category_id = ?", categoryID).
		Preload("Category").Preload("Uploader")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&documents).Error; err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}

func (r *documentRepository) FindPublic(page, limit int) ([]models.Document, int64, error) {
	var documents []models.Document
	var total int64

	query := r.db.Model(&models.Document{}).Where("is_public = ? AND status = ?", true, "active").
		Preload("Category").Preload("Uploader")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&documents).Error; err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}
