package repositories

import (
	"site-admin-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type EventFlyerRepository interface {
	FindAll(filters map[string]interface{}) ([]models.EventFlyer, error)
	FindByID(id uint) (*models.EventFlyer, error)
	FindActive() ([]models.EventFlyer, error)
	FindUpcoming(limit int) ([]models.EventFlyer, error)
	FindByDateRange(start, end time.Time) ([]models.EventFlyer, error)
	Create(flyer *models.EventFlyer) error
	Update(flyer *models.EventFlyer) error
	Delete(id uint) error
}

type eventFlyerRepository struct {
	db *gorm.DB
}

func NewEventFlyerRepository(db *gorm.DB) EventFlyerRepository {
	return &eventFlyerRepository{db: db}
}

func (r *eventFlyerRepository) FindAll(filters map[string]interface{}) ([]models.EventFlyer, error) {
	var flyers []models.EventFlyer
	query := r.db.Model(&models.EventFlyer{}).Preload("Creator")

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Order("event_date DESC, order_number ASC").Find(&flyers).Error; err != nil {
		return nil, err
	}
	return flyers, nil
}

func (r *eventFlyerRepository) FindByID(id uint) (*models.EventFlyer, error) {
	var flyer models.EventFlyer
	if err := r.db.Preload("Creator").First(&flyer, id).Error; err != nil {
		return nil, err
	}
	return &flyer, nil
}

func (r *eventFlyerRepository) FindActive() ([]models.EventFlyer, error) {
	var flyers []models.EventFlyer
	now := time.Now()
	if err := r.db.Where("is_active = ? AND (start_display_date IS NULL OR start_display_date <= ?) AND (end_display_date IS NULL OR end_display_date >= ?)",
		true, now, now).Preload("Creator").Order("order_number ASC").Find(&flyers).Error; err != nil {
		return nil, err
	}
	return flyers, nil
}

func (r *eventFlyerRepository) FindUpcoming(limit int) ([]models.EventFlyer, error) {
	var flyers []models.EventFlyer
	now := time.Now()
	if err := r.db.Where("is_active = ? AND event_date >= ?", true, now).
		Preload("Creator").Order("event_date ASC").Limit(limit).Find(&flyers).Error; err != nil {
		return nil, err
	}
	return flyers, nil
}

func (r *eventFlyerRepository) FindByDateRange(start, end time.Time) ([]models.EventFlyer, error) {
	var flyers []models.EventFlyer
	if err := r.db.Where("event_date BETWEEN ? AND ?", start, end).
		Preload("Creator").Order("event_date ASC").Find(&flyers).Error; err != nil {
		return nil, err
	}
	return flyers, nil
}

func (r *eventFlyerRepository) Create(flyer *models.EventFlyer) error {
	return r.db.Create(flyer).Error
}

func (r *eventFlyerRepository) Update(flyer *models.EventFlyer) error {
	return r.db.Save(flyer).Error
}

func (r *eventFlyerRepository) Delete(id uint) error {
	return r.db.Delete(&models.EventFlyer{}, id).Error
}
