package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type DepartmentService interface {
	GetAll(filters map[string]interface{}) ([]models.DepartmentResponse, error)
	GetByID(id uint) (*models.DepartmentResponse, error)
	Update(id uint, req *models.DepartmentUpdateRequest) (*models.DepartmentResponse, error)
}

type departmentService struct {
	departmentRepo repositories.DepartmentRepository
	cfg            *config.Config
}

func NewDepartmentService(departmentRepo repositories.DepartmentRepository, cfg *config.Config) DepartmentService {
	return &departmentService{
		departmentRepo: departmentRepo,
		cfg:            cfg,
	}
}

func (s *departmentService) GetAll(filters map[string]interface{}) ([]models.DepartmentResponse, error) {
	departments, err := s.departmentRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	var responses []models.DepartmentResponse
	for _, dept := range departments {
		responses = append(responses, s.toResponse(&dept))
	}

	return responses, nil
}

func (s *departmentService) GetByID(id uint) (*models.DepartmentResponse, error) {
	dept, err := s.departmentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(dept)
	return &response, nil
}

func (s *departmentService) Update(id uint, req *models.DepartmentUpdateRequest) (*models.DepartmentResponse, error) {
	dept, err := s.departmentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		dept.Name = *req.Name
	}
	if req.Description != nil {
		dept.Description = req.Description
	}
	if req.HeadName != nil {
		dept.HeadName = req.HeadName
	}
	if req.OrderNumber != nil {
		dept.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		dept.IsActive = *req.IsActive
	}

	if err := s.departmentRepo.Update(dept); err != nil {
		return nil, err
	}

	response := s.toResponse(dept)
	return &response, nil
}

func (s *departmentService) toResponse(dept *models.Department) models.DepartmentResponse {
	return models.DepartmentResponse{
		ID:          dept.ID,
		Name:        dept.Name,
		Description: dept.Description,
		HeadName:    dept.HeadName,
		OrderNumber: dept.OrderNumber,
		IsActive:    dept.IsActive,
		CreatedAt:   dept.CreatedAt,
		UpdatedAt:   dept.UpdatedAt,
	}
}
