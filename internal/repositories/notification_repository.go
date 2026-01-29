package repositories

import (
	"site-admin-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	FindByUserID(userID uint, page, limit int) ([]models.Notification, int64, error)
	FindByID(id uint) (*models.Notification, error)
	Create(notification *models.Notification) error
	MarkAsRead(id uint) error
	MarkAllAsRead(userID uint) error
	Delete(id uint) error
	GetUnreadCount(userID uint) (int64, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) FindByUserID(userID uint, page, limit int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64

	query := r.db.Model(&models.Notification{}).Where("user_id = ?", userID)

	// Count total
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&notifications).Error

	return notifications, total, err
}

func (r *notificationRepository) FindByID(id uint) (*models.Notification, error) {
	var notification models.Notification
	err := r.db.First(&notification, id).Error
	return &notification, err
}

func (r *notificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) MarkAsRead(id uint) error {
	now := time.Now()
	return r.db.Model(&models.Notification{}).Where("id = ?", id).Update("read_at", now).Error
}

func (r *notificationRepository) MarkAllAsRead(userID uint) error {
	now := time.Now()
	return r.db.Model(&models.Notification{}).Where("user_id = ? AND read_at IS NULL", userID).Update("read_at", now).Error
}

func (r *notificationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Notification{}, id).Error
}

func (r *notificationRepository) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Notification{}).Where("user_id = ? AND read_at IS NULL", userID).Count(&count).Error
	return count, err
}
