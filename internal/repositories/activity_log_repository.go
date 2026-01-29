package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type ActivityLogRepository interface {
	FindAll(page, limit int, filters map[string]interface{}) ([]models.ActivityLog, int64, error)
	FindByID(id uint) (*models.ActivityLog, error)
	Create(log *models.ActivityLog) error
}

type activityLogRepository struct {
	db *gorm.DB
}

func NewActivityLogRepository(db *gorm.DB) ActivityLogRepository {
	return &activityLogRepository{db: db}
}

func (r *activityLogRepository) FindAll(page, limit int, filters map[string]interface{}) ([]models.ActivityLog, int64, error) {
	var logs []models.ActivityLog
	var total int64

	query := r.db.Model(&models.ActivityLog{}).Preload("User")

	// Apply filters
	if userID, ok := filters["user_id"]; ok {
		query = query.Where("user_id = ?", userID)
	}
	if action, ok := filters["action"]; ok {
		query = query.Where("action = ?", action)
	}
	if entityType, ok := filters["entity_type"]; ok {
		query = query.Where("entity_type = ?", entityType)
	}

	// Count total
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&logs).Error

	return logs, total, err
}

func (r *activityLogRepository) FindByID(id uint) (*models.ActivityLog, error) {
	var log models.ActivityLog
	err := r.db.Preload("User").First(&log, id).Error
	return &log, err
}

func (r *activityLogRepository) Create(log *models.ActivityLog) error {
	return r.db.Create(log).Error
}
