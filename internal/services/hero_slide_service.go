package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type HeroSlideService interface {
	GetAll() ([]models.HeroSlideResponse, error)
	GetByID(id uint) (*models.HeroSlideResponse, error)
	Create(req *models.HeroSlideCreateRequest) (*models.HeroSlideResponse, error)
	Update(id uint, req *models.HeroSlideUpdateRequest) (*models.HeroSlideResponse, error)
	Delete(id uint) error
	Reorder(req *models.HeroSlideReorderRequest) error
	GetActive() ([]models.HeroSlideResponse, error)
}

type heroSlideService struct {
	heroSlideRepo repositories.HeroSlideRepository
	cfg           *config.Config
}

func NewHeroSlideService(heroSlideRepo repositories.HeroSlideRepository, cfg *config.Config) HeroSlideService {
	return &heroSlideService{
		heroSlideRepo: heroSlideRepo,
		cfg:           cfg,
	}
}

func (s *heroSlideService) GetAll() ([]models.HeroSlideResponse, error) {
	filters := make(map[string]interface{})
	slides, err := s.heroSlideRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	var responses []models.HeroSlideResponse
	for _, slide := range slides {
		responses = append(responses, s.toResponse(&slide))
	}

	return responses, nil
}

func (s *heroSlideService) GetByID(id uint) (*models.HeroSlideResponse, error) {
	slide, err := s.heroSlideRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(slide)
	return &response, nil
}

func (s *heroSlideService) Create(req *models.HeroSlideCreateRequest) (*models.HeroSlideResponse, error) {
	slide := &models.HeroSlide{
		Title:             req.Title,
		Description:       req.Description,
		Image:             req.Image,
		CTALabel:          req.CTALabel,
		CTAHref:           req.CTAHref,
		CTASecondaryLabel: req.CTASecondaryLabel,
		CTASecondaryHref:  req.CTASecondaryHref,
		OrderNumber:       req.OrderNumber,
		IsActive:          req.IsActive,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
	}

	if err := s.heroSlideRepo.Create(slide); err != nil {
		return nil, err
	}

	response := s.toResponse(slide)
	return &response, nil
}

func (s *heroSlideService) Update(id uint, req *models.HeroSlideUpdateRequest) (*models.HeroSlideResponse, error) {
	slide, err := s.heroSlideRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		slide.Title = *req.Title
	}
	if req.Description != nil {
		slide.Description = req.Description
	}
	if req.Image != nil {
		slide.Image = *req.Image
	}
	if req.CTALabel != nil {
		slide.CTALabel = req.CTALabel
	}
	if req.CTAHref != nil {
		slide.CTAHref = req.CTAHref
	}
	if req.CTASecondaryLabel != nil {
		slide.CTASecondaryLabel = req.CTASecondaryLabel
	}
	if req.CTASecondaryHref != nil {
		slide.CTASecondaryHref = req.CTASecondaryHref
	}
	if req.OrderNumber != nil {
		slide.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		slide.IsActive = *req.IsActive
	}
	if req.StartDate != nil {
		slide.StartDate = req.StartDate
	}
	if req.EndDate != nil {
		slide.EndDate = req.EndDate
	}

	if err := s.heroSlideRepo.Update(slide); err != nil {
		return nil, err
	}

	response := s.toResponse(slide)
	return &response, nil
}

func (s *heroSlideService) Delete(id uint) error {
	return s.heroSlideRepo.Delete(id)
}

func (s *heroSlideService) Reorder(req *models.HeroSlideReorderRequest) error {
	for _, order := range req.SlideOrders {
		slide, err := s.heroSlideRepo.FindByID(order.ID)
		if err != nil {
			return err
		}
		slide.OrderNumber = order.Order
		if err := s.heroSlideRepo.Update(slide); err != nil {
			return err
		}
	}
	return nil
}

func (s *heroSlideService) GetActive() ([]models.HeroSlideResponse, error) {
	filters := map[string]interface{}{"is_active": true}
	slides, err := s.heroSlideRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	var responses []models.HeroSlideResponse
	for _, slide := range slides {
		responses = append(responses, s.toResponse(&slide))
	}

	return responses, nil
}

func (s *heroSlideService) toResponse(slide *models.HeroSlide) models.HeroSlideResponse {
	return models.HeroSlideResponse{
		ID:                slide.ID,
		Title:             slide.Title,
		Subtitle:          slide.Description,
		Description:       slide.Description,
		Image:             slide.Image,
		ImageURL:          slide.Image,
		CTALabel:          slide.CTALabel,
		CTAHref:           slide.CTAHref,
		LinkURL:           slide.CTAHref,
		LinkText:          slide.CTALabel,
		CTASecondaryLabel: slide.CTASecondaryLabel,
		CTASecondaryHref:  slide.CTASecondaryHref,
		OrderNumber:       slide.OrderNumber,
		DisplayOrder:      slide.OrderNumber,
		IsActive:          slide.IsActive,
		StartDate:         slide.StartDate,
		EndDate:           slide.EndDate,
		CreatedAt:         slide.CreatedAt,
		UpdatedAt:         slide.UpdatedAt,
	}
}
