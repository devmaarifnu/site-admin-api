package services

import (
	"errors"
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
	"site-admin-api/internal/utils"

	"gorm.io/gorm"
)

// Category Service
type CategoryService interface {
	GetAll(filters map[string]interface{}) ([]models.CategoryResponse, error)
	GetByID(id uint) (*models.CategoryResponse, error)
	GetBySlug(slug string) (*models.CategoryResponse, error)
	GetByType(categoryType string) ([]models.CategoryResponse, error)
	Create(req *models.CategoryCreateRequest) (*models.CategoryResponse, error)
	Update(id uint, req *models.CategoryUpdateRequest) (*models.CategoryResponse, error)
	Delete(id uint) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
	cfg          *config.Config
}

func NewCategoryService(categoryRepo repositories.CategoryRepository, cfg *config.Config) CategoryService {
	return &categoryService{categoryRepo: categoryRepo, cfg: cfg}
}

func (s *categoryService) GetAll(filters map[string]interface{}) ([]models.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CategoryResponse, len(categories))
	for i, c := range categories {
		responses[i] = s.toResponse(&c)
	}

	return responses, nil
}

func (s *categoryService) GetByID(id uint) (*models.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	response := s.toResponse(category)
	return &response, nil
}

func (s *categoryService) GetBySlug(slug string) (*models.CategoryResponse, error) {
	category, err := s.categoryRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	response := s.toResponse(category)
	return &response, nil
}

func (s *categoryService) GetByType(categoryType string) ([]models.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindByType(categoryType)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CategoryResponse, len(categories))
	for i, c := range categories {
		responses[i] = s.toResponse(&c)
	}

	return responses, nil
}

func (s *categoryService) Create(req *models.CategoryCreateRequest) (*models.CategoryResponse, error) {
	slug := utils.GenerateSlug(req.Name)

	category := &models.Category{
		Name:        req.Name,
		Slug:        slug,
		Description: req.Description,
		Type:        req.Type,
		Icon:        req.Icon,
		Color:       req.Color,
		IsActive:    req.IsActive,
		OrderNumber: req.OrderNumber,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	response := s.toResponse(category)
	return &response, nil
}

func (s *categoryService) Update(id uint, req *models.CategoryUpdateRequest) (*models.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	if req.Name != nil {
		category.Name = *req.Name
		category.Slug = utils.GenerateSlug(*req.Name)
	}
	if req.Description != nil {
		category.Description = req.Description
	}
	if req.Type != nil {
		category.Type = *req.Type
	}
	if req.Icon != nil {
		category.Icon = req.Icon
	}
	if req.Color != nil {
		category.Color = req.Color
	}
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}
	if req.OrderNumber != nil {
		category.OrderNumber = *req.OrderNumber
	}

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	response := s.toResponse(category)
	return &response, nil
}

func (s *categoryService) Delete(id uint) error {
	_, err := s.categoryRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}

	return s.categoryRepo.Delete(id)
}

func (s *categoryService) toResponse(category *models.Category) models.CategoryResponse {
	return models.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
		Type:        category.Type,
		Icon:        category.Icon,
		Color:       category.Color,
		IsActive:    category.IsActive,
		OrderNumber: category.OrderNumber,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

// Tag Service
type TagService interface {
	GetAll() ([]models.TagResponse, error)
	GetByID(id uint) (*models.TagResponse, error)
	GetBySlug(slug string) (*models.TagResponse, error)
	Create(req *models.TagCreateRequest) (*models.TagResponse, error)
	Update(id uint, req *models.TagUpdateRequest) (*models.TagResponse, error)
	Delete(id uint) error
}

type tagService struct {
	tagRepo repositories.TagRepository
	cfg     *config.Config
}

func NewTagService(tagRepo repositories.TagRepository, cfg *config.Config) TagService {
	return &tagService{tagRepo: tagRepo, cfg: cfg}
}

func (s *tagService) GetAll() ([]models.TagResponse, error) {
	tags, err := s.tagRepo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.TagResponse, len(tags))
	for i, t := range tags {
		responses[i] = s.toResponse(&t)
	}

	return responses, nil
}

func (s *tagService) GetByID(id uint) (*models.TagResponse, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	response := s.toResponse(tag)
	return &response, nil
}

func (s *tagService) GetBySlug(slug string) (*models.TagResponse, error) {
	tag, err := s.tagRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	response := s.toResponse(tag)
	return &response, nil
}

func (s *tagService) Create(req *models.TagCreateRequest) (*models.TagResponse, error) {
	slug := utils.GenerateSlug(req.Name)

	tag := &models.Tag{
		Name: req.Name,
		Slug: slug,
	}

	if err := s.tagRepo.Create(tag); err != nil {
		return nil, err
	}

	response := s.toResponse(tag)
	return &response, nil
}

func (s *tagService) Update(id uint, req *models.TagUpdateRequest) (*models.TagResponse, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	if req.Name != nil {
		tag.Name = *req.Name
		tag.Slug = utils.GenerateSlug(*req.Name)
	}

	if err := s.tagRepo.Update(tag); err != nil {
		return nil, err
	}

	response := s.toResponse(tag)
	return &response, nil
}

func (s *tagService) Delete(id uint) error {
	_, err := s.tagRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("tag not found")
		}
		return err
	}

	return s.tagRepo.Delete(id)
}

func (s *tagService) toResponse(tag *models.Tag) models.TagResponse {
	return models.TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		Slug:      tag.Slug,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}
}
