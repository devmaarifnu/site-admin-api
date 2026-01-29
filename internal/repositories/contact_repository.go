package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type ContactMessageRepository interface {
	FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.ContactMessage, int64, error)
	FindByID(id uint) (*models.ContactMessage, error)
	FindByTicketID(ticketID string) (*models.ContactMessage, error)
	FindByStatus(status string, page, limit int) ([]models.ContactMessage, int64, error)
	FindByPriority(priority string, page, limit int) ([]models.ContactMessage, int64, error)
	FindAssignedTo(userID uint, page, limit int) ([]models.ContactMessage, int64, error)
	Create(message *models.ContactMessage) error
	Update(message *models.ContactMessage) error
	Delete(id uint) error
	GetNextTicketNumber() (int, error)
}

type contactMessageRepository struct {
	db *gorm.DB
}

func NewContactMessageRepository(db *gorm.DB) ContactMessageRepository {
	return &contactMessageRepository{db: db}
}

func (r *contactMessageRepository) FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.ContactMessage, int64, error) {
	var messages []models.ContactMessage
	var total int64

	query := r.db.Model(&models.ContactMessage{}).Preload("AssignedUser")

	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ? OR subject LIKE ? OR message LIKE ? OR ticket_id LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
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
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

func (r *contactMessageRepository) FindByID(id uint) (*models.ContactMessage, error) {
	var message models.ContactMessage
	if err := r.db.Preload("AssignedUser").First(&message, id).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *contactMessageRepository) FindByTicketID(ticketID string) (*models.ContactMessage, error) {
	var message models.ContactMessage
	if err := r.db.Preload("AssignedUser").Where("ticket_id = ?", ticketID).First(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *contactMessageRepository) FindByStatus(status string, page, limit int) ([]models.ContactMessage, int64, error) {
	var messages []models.ContactMessage
	var total int64

	query := r.db.Model(&models.ContactMessage{}).Where("status = ?", status).Preload("AssignedUser")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

func (r *contactMessageRepository) FindByPriority(priority string, page, limit int) ([]models.ContactMessage, int64, error) {
	var messages []models.ContactMessage
	var total int64

	query := r.db.Model(&models.ContactMessage{}).Where("priority = ?", priority).Preload("AssignedUser")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

func (r *contactMessageRepository) FindAssignedTo(userID uint, page, limit int) ([]models.ContactMessage, int64, error) {
	var messages []models.ContactMessage
	var total int64

	query := r.db.Model(&models.ContactMessage{}).Where("assigned_to = ?", userID).Preload("AssignedUser")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

func (r *contactMessageRepository) Create(message *models.ContactMessage) error {
	return r.db.Create(message).Error
}

func (r *contactMessageRepository) Update(message *models.ContactMessage) error {
	return r.db.Save(message).Error
}

func (r *contactMessageRepository) Delete(id uint) error {
	return r.db.Delete(&models.ContactMessage{}, id).Error
}

func (r *contactMessageRepository) GetNextTicketNumber() (int, error) {
	var count int64
	if err := r.db.Model(&models.ContactMessage{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count) + 1, nil
}
