package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type PengurusService interface {
	GetAll(filters map[string]interface{}) ([]models.PengurusResponse, error)
	GetByID(id uint) (*models.PengurusResponse, error)
	Create(req *models.PengurusCreateRequest) (*models.PengurusResponse, error)
	Update(id uint, req *models.PengurusUpdateRequest) (*models.PengurusResponse, error)
	Delete(id uint) error
	Reorder(orders []map[string]interface{}) error
}

type pengurusService struct {
	pengurusRepo repositories.PengurusRepository
	cfg          *config.Config
}

func NewPengurusService(pengurusRepo repositories.PengurusRepository, cfg *config.Config) PengurusService {
	return &pengurusService{
		pengurusRepo: pengurusRepo,
		cfg:          cfg,
	}
}

func (s *pengurusService) GetAll(filters map[string]interface{}) ([]models.PengurusResponse, error) {
	pengurusList, err := s.pengurusRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	var responses []models.PengurusResponse
	for _, pengurus := range pengurusList {
		responses = append(responses, s.toResponse(&pengurus))
	}

	return responses, nil
}

func (s *pengurusService) GetByID(id uint) (*models.PengurusResponse, error) {
	pengurus, err := s.pengurusRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(pengurus)
	return &response, nil
}

func (s *pengurusService) Create(req *models.PengurusCreateRequest) (*models.PengurusResponse, error) {
	pengurus := &models.Pengurus{
		Nama:           req.Nama,
		Jabatan:        req.Jabatan,
		Kategori:       req.Kategori,
		Foto:           req.Foto,
		Bio:            req.Bio,
		Email:          req.Email,
		Phone:          req.Phone,
		PeriodeMulai:   req.PeriodeMulai,
		PeriodeSelesai: req.PeriodeSelesai,
		OrderNumber:    req.OrderNumber,
		IsActive:       req.IsActive,
	}

	if err := s.pengurusRepo.Create(pengurus); err != nil {
		return nil, err
	}

	response := s.toResponse(pengurus)
	return &response, nil
}

func (s *pengurusService) Update(id uint, req *models.PengurusUpdateRequest) (*models.PengurusResponse, error) {
	pengurus, err := s.pengurusRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Nama != nil {
		pengurus.Nama = *req.Nama
	}
	if req.Jabatan != nil {
		pengurus.Jabatan = *req.Jabatan
	}
	if req.Kategori != nil {
		pengurus.Kategori = *req.Kategori
	}
	if req.Foto != nil {
		pengurus.Foto = req.Foto
	}
	if req.Bio != nil {
		pengurus.Bio = req.Bio
	}
	if req.Email != nil {
		pengurus.Email = req.Email
	}
	if req.Phone != nil {
		pengurus.Phone = req.Phone
	}
	if req.PeriodeMulai != nil {
		pengurus.PeriodeMulai = *req.PeriodeMulai
	}
	if req.PeriodeSelesai != nil {
		pengurus.PeriodeSelesai = *req.PeriodeSelesai
	}
	if req.OrderNumber != nil {
		pengurus.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		pengurus.IsActive = *req.IsActive
	}

	if err := s.pengurusRepo.Update(pengurus); err != nil {
		return nil, err
	}

	response := s.toResponse(pengurus)
	return &response, nil
}

func (s *pengurusService) Delete(id uint) error {
	return s.pengurusRepo.Delete(id)
}

func (s *pengurusService) Reorder(orders []map[string]interface{}) error {
	for _, order := range orders {
		id := uint(order["id"].(float64))
		orderNum := int(order["order_number"].(float64))

		pengurus, err := s.pengurusRepo.FindByID(id)
		if err != nil {
			continue
		}

		pengurus.OrderNumber = orderNum
		s.pengurusRepo.Update(pengurus)
	}

	return nil
}

func (s *pengurusService) toResponse(pengurus *models.Pengurus) models.PengurusResponse {
	return models.PengurusResponse{
		ID:             pengurus.ID,
		Nama:           pengurus.Nama,
		Jabatan:        pengurus.Jabatan,
		Kategori:       pengurus.Kategori,
		Foto:           pengurus.Foto,
		Bio:            pengurus.Bio,
		Email:          pengurus.Email,
		Phone:          pengurus.Phone,
		PeriodeMulai:   pengurus.PeriodeMulai,
		PeriodeSelesai: pengurus.PeriodeSelesai,
		OrderNumber:    pengurus.OrderNumber,
		IsActive:       pengurus.IsActive,
		CreatedAt:      pengurus.CreatedAt,
		UpdatedAt:      pengurus.UpdatedAt,
	}
}
