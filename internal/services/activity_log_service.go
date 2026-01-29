package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type ActivityLogService interface {
	GetAll(page, limit int, filters map[string]interface{}) ([]models.ActivityLogResponse, int64, error)
	GetByID(id uint) (*models.ActivityLogResponse, error)
}

type activityLogService struct {
	activityLogRepo repositories.ActivityLogRepository
	cfg             *config.Config
}

func NewActivityLogService(activityLogRepo repositories.ActivityLogRepository, cfg *config.Config) ActivityLogService {
	return &activityLogService{
		activityLogRepo: activityLogRepo,
		cfg:             cfg,
	}
}

func (s *activityLogService) GetAll(page, limit int, filters map[string]interface{}) ([]models.ActivityLogResponse, int64, error) {
	logs, total, err := s.activityLogRepo.FindAll(page, limit, filters)
	if err != nil {
		return nil, 0, err
	}

	var responses []models.ActivityLogResponse
	for _, log := range logs {
		responses = append(responses, s.toResponse(&log))
	}

	return responses, total, nil
}

func (s *activityLogService) GetByID(id uint) (*models.ActivityLogResponse, error) {
	log, err := s.activityLogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(log)
	return &response, nil
}

func (s *activityLogService) toResponse(log *models.ActivityLog) models.ActivityLogResponse {
	return models.ActivityLogResponse{
		ID:          log.ID,
		LogName:     log.LogName,
		Description: log.Description,
		SubjectType: log.SubjectType,
		SubjectID:   log.SubjectID,
		CauserType:  log.CauserType,
		CauserID:    log.CauserID,
		Properties:  log.Properties,
		IPAddress:   log.IPAddress,
		UserAgent:   log.UserAgent,
		CreatedAt:   log.CreatedAt,
	}
}
