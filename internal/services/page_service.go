package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type PageService interface {
	GetAll() ([]models.PageResponse, error)
	GetByID(id uint) (*models.PageResponse, error)
	GetBySlug(slug string) (*models.PageResponse, error)
	Update(id uint, req *models.PageUpdateRequest) (*models.PageResponse, error)
}

type pageService struct {
	pageRepo repositories.PageRepository
	cfg      *config.Config
}

func NewPageService(pageRepo repositories.PageRepository, cfg *config.Config) PageService {
	return &pageService{
		pageRepo: pageRepo,
		cfg:      cfg,
	}
}

func (s *pageService) GetAll() ([]models.PageResponse, error) {
	pages, err := s.pageRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.PageResponse
	for _, page := range pages {
		responses = append(responses, s.toResponse(&page))
	}

	return responses, nil
}

func (s *pageService) GetByID(id uint) (*models.PageResponse, error) {
	page, err := s.pageRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(page)
	return &response, nil
}

func (s *pageService) GetBySlug(slug string) (*models.PageResponse, error) {
	page, err := s.pageRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(page)
	return &response, nil
}

func (s *pageService) Update(id uint, req *models.PageUpdateRequest) (*models.PageResponse, error) {
	page, err := s.pageRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		page.Title = *req.Title
	}
	if req.Content != nil {
		page.Content = *req.Content
	}
	if req.MetaTitle != nil {
		page.MetaTitle = req.MetaTitle
	}
	if req.MetaDescription != nil {
		page.MetaDescription = req.MetaDescription
	}
	if req.MetaKeywords != nil {
		page.MetaKeywords = req.MetaKeywords
	}

	if err := s.pageRepo.Update(page); err != nil {
		return nil, err
	}

	response := s.toResponse(page)
	return &response, nil
}

func (s *pageService) toResponse(page *models.Page) models.PageResponse {
	return models.PageResponse{
		ID:              page.ID,
		Title:           page.Title,
		Slug:            page.Slug,
		Content:         page.Content,
		MetaTitle:       page.MetaTitle,
		MetaDescription: page.MetaDescription,
		MetaKeywords:    page.MetaKeywords,
		CreatedAt:       page.CreatedAt,
		UpdatedAt:       page.UpdatedAt,
	}
}
