package services

import (
	"time"

	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type ContactMessageService interface {
	GetAll(page, limit int, filters map[string]interface{}) ([]models.ContactMessageResponse, int64, error)
	GetByID(id uint) (*models.ContactMessageResponse, error)
	UpdateStatus(id uint, status string) (*models.ContactMessageResponse, error)
	Reply(id uint, replyMessage string, replierID uint) (*models.ContactMessageResponse, error)
	Delete(id uint) error
}

type contactMessageService struct {
	contactRepo repositories.ContactMessageRepository
	cfg         *config.Config
}

func NewContactMessageService(contactRepo repositories.ContactMessageRepository, cfg *config.Config) ContactMessageService {
	return &contactMessageService{
		contactRepo: contactRepo,
		cfg:         cfg,
	}
}

func (s *contactMessageService) GetAll(page, limit int, filters map[string]interface{}) ([]models.ContactMessageResponse, int64, error) {
	messages, total, err := s.contactRepo.FindAll(page, limit, "", filters)
	if err != nil {
		return nil, 0, err
	}

	var responses []models.ContactMessageResponse
	for _, msg := range messages {
		responses = append(responses, s.toResponse(&msg))
	}

	return responses, total, nil
}

func (s *contactMessageService) GetByID(id uint) (*models.ContactMessageResponse, error) {
	message, err := s.contactRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(message)
	return &response, nil
}

func (s *contactMessageService) UpdateStatus(id uint, status string) (*models.ContactMessageResponse, error) {
	message, err := s.contactRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	message.Status = status
	if err := s.contactRepo.Update(message); err != nil {
		return nil, err
	}

	response := s.toResponse(message)
	return &response, nil
}

func (s *contactMessageService) Reply(id uint, replyMessage string, replierID uint) (*models.ContactMessageResponse, error) {
	message, err := s.contactRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update status and replied time
	message.Status = "replied"
	now := time.Now()
	message.RepliedAt = &now

	// Save notes as reply
	message.Notes = &replyMessage

	if err := s.contactRepo.Update(message); err != nil {
		return nil, err
	}

	response := s.toResponse(message)
	return &response, nil
}

func (s *contactMessageService) Delete(id uint) error {
	return s.contactRepo.Delete(id)
}

func (s *contactMessageService) toResponse(msg *models.ContactMessage) models.ContactMessageResponse {
	return models.ContactMessageResponse{
		ID:           msg.ID,
		TicketID:     msg.TicketID,
		Name:         msg.Name,
		Email:        msg.Email,
		Phone:        msg.Phone,
		Subject:      msg.Subject,
		Message:      msg.Message,
		Status:       msg.Status,
		Priority:     msg.Priority,
		IPAddress:    msg.IPAddress,
		UserAgent:    msg.UserAgent,
		AssignedTo:   msg.AssignedTo,
		AssignedUser: msg.AssignedUser,
		RepliedAt:    msg.RepliedAt,
		ResolvedAt:   msg.ResolvedAt,
		Notes:        msg.Notes,
		CreatedAt:    msg.CreatedAt,
		UpdatedAt:    msg.UpdatedAt,
	}
}
