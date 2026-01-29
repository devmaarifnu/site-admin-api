package services

import (
	"errors"
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
	"site-admin-api/internal/utils"
	"time"

	"gorm.io/gorm"
)

type OpinionService interface {
	GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.OpinionArticleResponse, int64, error)
	GetByID(id uint) (*models.OpinionArticleResponse, error)
	GetBySlug(slug string) (*models.OpinionArticleResponse, error)
	Create(req *models.OpinionArticleCreateRequest, createdBy uint) (*models.OpinionArticleResponse, error)
	Update(id uint, req *models.OpinionArticleUpdateRequest) (*models.OpinionArticleResponse, error)
	Delete(id uint) error
	IncrementViews(id uint) error
	GetFeatured(limit int) ([]models.OpinionArticleResponse, error)
	GetByTag(tagID uint, page, limit int) ([]models.OpinionArticleResponse, int64, error)
}

type opinionService struct {
	opinionRepo repositories.OpinionRepository
	cfg         *config.Config
}

func NewOpinionService(opinionRepo repositories.OpinionRepository, cfg *config.Config) OpinionService {
	return &opinionService{opinionRepo: opinionRepo, cfg: cfg}
}

func (s *opinionService) GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.OpinionArticleResponse, int64, error) {
	opinions, total, err := s.opinionRepo.FindAll(page, limit, search, filters)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.OpinionArticleResponse, len(opinions))
	for i, o := range opinions {
		responses[i] = s.toResponse(&o)
	}

	return responses, total, nil
}

func (s *opinionService) GetByID(id uint) (*models.OpinionArticleResponse, error) {
	opinion, err := s.opinionRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("opinion article not found")
		}
		return nil, err
	}

	response := s.toResponse(opinion)
	return &response, nil
}

func (s *opinionService) GetBySlug(slug string) (*models.OpinionArticleResponse, error) {
	opinion, err := s.opinionRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("opinion article not found")
		}
		return nil, err
	}

	response := s.toResponse(opinion)
	return &response, nil
}

func (s *opinionService) Create(req *models.OpinionArticleCreateRequest, createdBy uint) (*models.OpinionArticleResponse, error) {
	slug := utils.GenerateSlug(req.Title)

	_, err := s.opinionRepo.FindBySlug(slug)
	if err == nil {
		slug = slug + "-" + time.Now().Format("20060102150405")
	}

	if req.Status == "published" && req.PublishedAt == nil {
		now := time.Now()
		req.PublishedAt = &now
	}

	opinion := &models.OpinionArticle{
		Title:           req.Title,
		Slug:            slug,
		Excerpt:         req.Excerpt,
		Content:         req.Content,
		Image:           req.Image,
		AuthorName:      req.AuthorName,
		AuthorTitle:     req.AuthorTitle,
		AuthorImage:     req.AuthorImage,
		AuthorBio:       req.AuthorBio,
		Status:          req.Status,
		PublishedAt:     req.PublishedAt,
		IsFeatured:      req.IsFeatured,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
		MetaKeywords:    req.MetaKeywords,
		CreatedBy:       &createdBy,
	}

	if err := s.opinionRepo.Create(opinion); err != nil {
		return nil, err
	}

	if len(req.TagIDs) > 0 {
		if err := s.opinionRepo.UpdateTags(opinion.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	opinion, err = s.opinionRepo.FindByID(opinion.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(opinion)
	return &response, nil
}

func (s *opinionService) Update(id uint, req *models.OpinionArticleUpdateRequest) (*models.OpinionArticleResponse, error) {
	opinion, err := s.opinionRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("opinion article not found")
		}
		return nil, err
	}

	if req.Title != nil {
		opinion.Title = *req.Title
		opinion.Slug = utils.GenerateSlug(*req.Title)
	}
	if req.Excerpt != nil {
		opinion.Excerpt = *req.Excerpt
	}
	if req.Content != nil {
		opinion.Content = *req.Content
	}
	if req.Image != nil {
		opinion.Image = req.Image
	}
	if req.AuthorName != nil {
		opinion.AuthorName = *req.AuthorName
	}
	if req.AuthorTitle != nil {
		opinion.AuthorTitle = req.AuthorTitle
	}
	if req.AuthorImage != nil {
		opinion.AuthorImage = req.AuthorImage
	}
	if req.AuthorBio != nil {
		opinion.AuthorBio = req.AuthorBio
	}
	if req.Status != nil {
		opinion.Status = *req.Status
		if *req.Status == "published" && opinion.PublishedAt == nil {
			now := time.Now()
			opinion.PublishedAt = &now
		}
	}
	if req.PublishedAt != nil {
		opinion.PublishedAt = req.PublishedAt
	}
	if req.IsFeatured != nil {
		opinion.IsFeatured = *req.IsFeatured
	}
	if req.MetaTitle != nil {
		opinion.MetaTitle = req.MetaTitle
	}
	if req.MetaDescription != nil {
		opinion.MetaDescription = req.MetaDescription
	}
	if req.MetaKeywords != nil {
		opinion.MetaKeywords = req.MetaKeywords
	}

	if err := s.opinionRepo.Update(opinion); err != nil {
		return nil, err
	}

	if req.TagIDs != nil {
		if err := s.opinionRepo.UpdateTags(opinion.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	opinion, err = s.opinionRepo.FindByID(opinion.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(opinion)
	return &response, nil
}

func (s *opinionService) Delete(id uint) error {
	_, err := s.opinionRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("opinion article not found")
		}
		return err
	}

	return s.opinionRepo.Delete(id)
}

func (s *opinionService) IncrementViews(id uint) error {
	return s.opinionRepo.IncrementViews(id)
}

func (s *opinionService) GetFeatured(limit int) ([]models.OpinionArticleResponse, error) {
	opinions, err := s.opinionRepo.FindFeatured(limit)
	if err != nil {
		return nil, err
	}

	responses := make([]models.OpinionArticleResponse, len(opinions))
	for i, o := range opinions {
		responses[i] = s.toResponse(&o)
	}

	return responses, nil
}

func (s *opinionService) GetByTag(tagID uint, page, limit int) ([]models.OpinionArticleResponse, int64, error) {
	opinions, total, err := s.opinionRepo.FindByTag(tagID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.OpinionArticleResponse, len(opinions))
	for i, o := range opinions {
		responses[i] = s.toResponse(&o)
	}

	return responses, total, nil
}

func (s *opinionService) toResponse(opinion *models.OpinionArticle) models.OpinionArticleResponse {
	return models.OpinionArticleResponse{
		ID:              opinion.ID,
		Title:           opinion.Title,
		Slug:            opinion.Slug,
		Excerpt:         opinion.Excerpt,
		Content:         opinion.Content,
		Image:           opinion.Image,
		AuthorName:      opinion.AuthorName,
		AuthorTitle:     opinion.AuthorTitle,
		AuthorImage:     opinion.AuthorImage,
		AuthorBio:       opinion.AuthorBio,
		Status:          opinion.Status,
		PublishedAt:     opinion.PublishedAt,
		Views:           opinion.Views,
		IsFeatured:      opinion.IsFeatured,
		MetaTitle:       opinion.MetaTitle,
		MetaDescription: opinion.MetaDescription,
		MetaKeywords:    opinion.MetaKeywords,
		CreatedBy:       opinion.CreatedBy,
		Creator:         opinion.Creator,
		Tags:            opinion.Tags,
		CreatedAt:       opinion.CreatedAt,
		UpdatedAt:       opinion.UpdatedAt,
	}
}
