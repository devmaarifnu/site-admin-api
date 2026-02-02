package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type SettingService interface {
	GetAll(group string) ([]models.SettingResponse, error)
	GetByKey(key string) (*models.SettingResponse, error)
	Update(key string, value string) (*models.SettingResponse, error)
	BulkUpdate(settings []models.SettingUpdateRequest) error
}

type settingService struct {
	settingRepo repositories.SettingRepository
	cfg         *config.Config
}

func NewSettingService(settingRepo repositories.SettingRepository, cfg *config.Config) SettingService {
	return &settingService{
		settingRepo: settingRepo,
		cfg:         cfg,
	}
}

func (s *settingService) GetAll(group string) ([]models.SettingResponse, error) {
	settings, err := s.settingRepo.FindAll(group)
	if err != nil {
		return nil, err
	}

	var responses []models.SettingResponse
	for _, setting := range settings {
		responses = append(responses, s.toResponse(&setting))
	}

	return responses, nil
}

func (s *settingService) GetByKey(key string) (*models.SettingResponse, error) {
	setting, err := s.settingRepo.FindByKey(key)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(setting)
	return &response, nil
}

func (s *settingService) Update(key string, value string) (*models.SettingResponse, error) {
	setting, err := s.settingRepo.FindByKey(key)
	if err != nil {
		return nil, err
	}

	setting.SettingValue = &value
	if err := s.settingRepo.Update(setting); err != nil {
		return nil, err
	}

	response := s.toResponse(setting)
	return &response, nil
}

func (s *settingService) BulkUpdate(settings []models.SettingUpdateRequest) error {
	for _, settingReq := range settings {
		// Try to find existing setting
		setting, err := s.settingRepo.FindByKey(settingReq.SettingKey)

		if err != nil {
			// Setting doesn't exist, create new one
			setting = &models.Setting{
				SettingKey:   settingReq.SettingKey,
				SettingValue: settingReq.SettingValue,
				SettingType:  "string",
				SettingGroup: "general",
				IsPublic:     false,
			}
		} else {
			// Update existing setting value
			if settingReq.SettingValue != nil {
				setting.SettingValue = settingReq.SettingValue
			}
		}

		// Apply optional fields
		if settingReq.SettingType != nil {
			setting.SettingType = *settingReq.SettingType
		}
		if settingReq.SettingGroup != nil {
			setting.SettingGroup = *settingReq.SettingGroup
		}
		if settingReq.Description != nil {
			setting.Description = settingReq.Description
		}
		if settingReq.IsPublic != nil {
			setting.IsPublic = *settingReq.IsPublic
		}

		// Upsert (create or update)
		if err := s.settingRepo.Upsert(setting); err != nil {
			return err
		}
	}
	return nil
}

func (s *settingService) toResponse(setting *models.Setting) models.SettingResponse {
	return models.SettingResponse{
		ID:           setting.ID,
		SettingKey:   setting.SettingKey,
		SettingValue: setting.SettingValue,
		SettingGroup: setting.SettingGroup,
		SettingType:  setting.SettingType,
		Description:  setting.Description,
		IsPublic:     setting.IsPublic,
		CreatedAt:    setting.CreatedAt,
		UpdatedAt:    setting.UpdatedAt,
	}
}
