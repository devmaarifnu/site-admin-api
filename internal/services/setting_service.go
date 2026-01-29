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
	if err := s.settingRepo.Update(key, value); err != nil {
		return nil, err
	}

	setting, err := s.settingRepo.FindByKey(key)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(setting)
	return &response, nil
}

func (s *settingService) BulkUpdate(settings []models.SettingUpdateRequest) error {
	for _, setting := range settings {
		if err := s.settingRepo.Update(setting.Key, setting.Value); err != nil {
			return err
		}
	}
	return nil
}

func (s *settingService) toResponse(setting *models.Setting) models.SettingResponse {
	return models.SettingResponse{
		ID:          setting.ID,
		Key:         setting.Key,
		Value:       setting.Value,
		Group:       setting.Group,
		Type:        setting.Type,
		Description: setting.Description,
		UpdatedAt:   setting.UpdatedAt,
	}
}
