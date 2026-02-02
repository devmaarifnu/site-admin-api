package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type EditorialTeamService interface {
	GetAll(filters map[string]interface{}) ([]models.EditorialTeamResponse, error)
	GetByID(id uint) (*models.EditorialTeamResponse, error)
	Update(id uint, req *models.EditorialTeamUpdateRequest) (*models.EditorialTeamResponse, error)
}

type editorialTeamService struct {
	editorialTeamRepo repositories.EditorialTeamRepository
	cfg               *config.Config
}

func NewEditorialTeamService(editorialTeamRepo repositories.EditorialTeamRepository, cfg *config.Config) EditorialTeamService {
	return &editorialTeamService{
		editorialTeamRepo: editorialTeamRepo,
		cfg:               cfg,
	}
}

func (s *editorialTeamService) GetAll(filters map[string]interface{}) ([]models.EditorialTeamResponse, error) {
	members, err := s.editorialTeamRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	var responses []models.EditorialTeamResponse
	for _, member := range members {
		responses = append(responses, s.toResponse(&member))
	}

	return responses, nil
}

func (s *editorialTeamService) GetByID(id uint) (*models.EditorialTeamResponse, error) {
	member, err := s.editorialTeamRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(member)
	return &response, nil
}

func (s *editorialTeamService) Update(id uint, req *models.EditorialTeamUpdateRequest) (*models.EditorialTeamResponse, error) {
	member, err := s.editorialTeamRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		member.Name = *req.Name
	}
	if req.Position != nil {
		member.Position = *req.Position
	}
	if req.RoleType != nil {
		member.RoleType = *req.RoleType
	}
	if req.Photo != nil {
		member.Photo = req.Photo
	}
	if req.Bio != nil {
		member.Bio = req.Bio
	}
	if req.Email != nil {
		member.Email = req.Email
	}
	if req.Phone != nil {
		member.Phone = req.Phone
	}
	if req.OrderNumber != nil {
		member.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		member.IsActive = *req.IsActive
	}

	if err := s.editorialTeamRepo.Update(member); err != nil {
		return nil, err
	}

	response := s.toResponse(member)
	return &response, nil
}

func (s *editorialTeamService) toResponse(member *models.EditorialTeam) models.EditorialTeamResponse {
	return models.EditorialTeamResponse{
		ID:          member.ID,
		Name:        member.Name,
		Position:    member.Position,
		RoleType:    member.RoleType,
		Photo:       member.Photo,
		Bio:         member.Bio,
		Email:       member.Email,
		Phone:       member.Phone,
		OrderNumber: member.OrderNumber,
		IsActive:    member.IsActive,
		CreatedAt:   member.CreatedAt,
		UpdatedAt:   member.UpdatedAt,
	}
}

// EditorialCouncil Service
type EditorialCouncilService interface {
	GetAll(filters map[string]interface{}) ([]models.EditorialCouncilResponse, error)
	GetByID(id uint) (*models.EditorialCouncilResponse, error)
	Update(id uint, req *models.EditorialCouncilUpdateRequest) (*models.EditorialCouncilResponse, error)
}

type editorialCouncilService struct {
	editorialCouncilRepo repositories.EditorialCouncilRepository
	cfg                  *config.Config
}

func NewEditorialCouncilService(editorialCouncilRepo repositories.EditorialCouncilRepository, cfg *config.Config) EditorialCouncilService {
	return &editorialCouncilService{
		editorialCouncilRepo: editorialCouncilRepo,
		cfg:                  cfg,
	}
}

func (s *editorialCouncilService) GetAll(filters map[string]interface{}) ([]models.EditorialCouncilResponse, error) {
	members, err := s.editorialCouncilRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	var responses []models.EditorialCouncilResponse
	for _, member := range members {
		responses = append(responses, s.toResponse(&member))
	}

	return responses, nil
}

func (s *editorialCouncilService) GetByID(id uint) (*models.EditorialCouncilResponse, error) {
	member, err := s.editorialCouncilRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(member)
	return &response, nil
}

func (s *editorialCouncilService) Update(id uint, req *models.EditorialCouncilUpdateRequest) (*models.EditorialCouncilResponse, error) {
	member, err := s.editorialCouncilRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		member.Name = *req.Name
	}
	if req.Institution != nil {
		member.Institution = *req.Institution
	}
	if req.Expertise != nil {
		member.Expertise = req.Expertise
	}
	if req.Photo != nil {
		member.Photo = req.Photo
	}
	if req.Bio != nil {
		member.Bio = req.Bio
	}
	if req.Email != nil {
		member.Email = req.Email
	}
	if req.OrderNumber != nil {
		member.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		member.IsActive = *req.IsActive
	}

	if err := s.editorialCouncilRepo.Update(member); err != nil {
		return nil, err
	}

	response := s.toResponse(member)
	return &response, nil
}

func (s *editorialCouncilService) toResponse(member *models.EditorialCouncil) models.EditorialCouncilResponse {
	return models.EditorialCouncilResponse{
		ID:          member.ID,
		Name:        member.Name,
		Institution: member.Institution,
		Expertise:   member.Expertise,
		Photo:       member.Photo,
		Bio:         member.Bio,
		Email:       member.Email,
		OrderNumber: member.OrderNumber,
		IsActive:    member.IsActive,
		CreatedAt:   member.CreatedAt,
		UpdatedAt:   member.UpdatedAt,
	}
}
