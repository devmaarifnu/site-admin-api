package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type NotificationService interface {
	GetAll(userID uint, page, limit int) ([]models.NotificationResponse, int64, error)
	GetByID(id uint) (*models.NotificationResponse, error)
	MarkAsRead(id uint) (*models.NotificationResponse, error)
	MarkAllAsRead(userID uint) error
	Delete(id uint) error
}

type notificationService struct {
	notificationRepo repositories.NotificationRepository
	cfg              *config.Config
}

func NewNotificationService(notificationRepo repositories.NotificationRepository, cfg *config.Config) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		cfg:              cfg,
	}
}

func (s *notificationService) GetAll(userID uint, page, limit int) ([]models.NotificationResponse, int64, error) {
	notifications, total, err := s.notificationRepo.FindByUserID(userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []models.NotificationResponse
	for _, notif := range notifications {
		responses = append(responses, s.toResponse(&notif))
	}

	return responses, total, nil
}

func (s *notificationService) GetByID(id uint) (*models.NotificationResponse, error) {
	notification, err := s.notificationRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(notification)
	return &response, nil
}

func (s *notificationService) MarkAsRead(id uint) (*models.NotificationResponse, error) {
	if err := s.notificationRepo.MarkAsRead(id); err != nil {
		return nil, err
	}

	notification, err := s.notificationRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(notification)
	return &response, nil
}

func (s *notificationService) MarkAllAsRead(userID uint) error {
	return s.notificationRepo.MarkAllAsRead(userID)
}

func (s *notificationService) Delete(id uint) error {
	return s.notificationRepo.Delete(id)
}

func (s *notificationService) toResponse(notif *models.Notification) models.NotificationResponse {
	return models.NotificationResponse{
		ID:        notif.ID,
		UserID:    notif.UserID,
		Type:      notif.Type,
		Title:     notif.Title,
		Message:   notif.Message,
		Data:      notif.Data,
		ReadAt:    notif.ReadAt,
		CreatedAt: notif.CreatedAt,
	}
}
