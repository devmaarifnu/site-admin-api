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

type NewsService interface {
	GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.NewsArticleResponse, int64, error)
	GetByID(id uint) (*models.NewsArticleResponse, error)
	GetBySlug(slug string) (*models.NewsArticleResponse, error)
	Create(req *models.NewsArticleCreateRequest, authorID uint) (*models.NewsArticleResponse, error)
	Update(id uint, req *models.NewsArticleUpdateRequest) (*models.NewsArticleResponse, error)
	Delete(id uint) error
	IncrementViews(id uint) error
	GetFeatured(limit int) ([]models.NewsArticleResponse, error)
	GetByCategory(categoryID uint, page, limit int) ([]models.NewsArticleResponse, int64, error)
	GetByTag(tagID uint, page, limit int) ([]models.NewsArticleResponse, int64, error)
}

type newsService struct {
	newsRepo repositories.NewsRepository
	cfg      *config.Config
}

func NewNewsService(newsRepo repositories.NewsRepository, cfg *config.Config) NewsService {
	return &newsService{newsRepo: newsRepo, cfg: cfg}
}

func (s *newsService) GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.NewsArticleResponse, int64, error) {
	news, total, err := s.newsRepo.FindAll(page, limit, search, filters)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.NewsArticleResponse, len(news))
	for i, n := range news {
		responses[i] = s.toResponse(&n)
	}

	return responses, total, nil
}

func (s *newsService) GetByID(id uint) (*models.NewsArticleResponse, error) {
	news, err := s.newsRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("news article not found")
		}
		return nil, err
	}

	response := s.toResponse(news)
	return &response, nil
}

func (s *newsService) GetBySlug(slug string) (*models.NewsArticleResponse, error) {
	news, err := s.newsRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("news article not found")
		}
		return nil, err
	}

	response := s.toResponse(news)
	return &response, nil
}

func (s *newsService) Create(req *models.NewsArticleCreateRequest, authorID uint) (*models.NewsArticleResponse, error) {
	// Generate slug
	slug := utils.GenerateSlug(req.Title)

	// Check if slug exists
	_, err := s.newsRepo.FindBySlug(slug)
	if err == nil {
		// Slug exists, append timestamp
		slug = slug + "-" + time.Now().Format("20060102150405")
	}

	// Set published_at if status is published and not provided
	if req.Status == "published" && req.PublishedAt == nil {
		now := time.Now()
		req.PublishedAt = &now
	}

	news := &models.NewsArticle{
		Title:           req.Title,
		Slug:            slug,
		Excerpt:         req.Excerpt,
		Content:         req.Content,
		Image:           req.Image,
		CategoryID:      req.CategoryID,
		AuthorID:        &authorID,
		Status:          req.Status,
		PublishedAt:     req.PublishedAt,
		IsFeatured:      req.IsFeatured,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
		MetaKeywords:    req.MetaKeywords,
	}

	if err := s.newsRepo.Create(news); err != nil {
		return nil, err
	}

	// Update tags
	if len(req.TagIDs) > 0 {
		if err := s.newsRepo.UpdateTags(news.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	// Reload with relationships
	news, err = s.newsRepo.FindByID(news.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(news)
	return &response, nil
}

func (s *newsService) Update(id uint, req *models.NewsArticleUpdateRequest) (*models.NewsArticleResponse, error) {
	news, err := s.newsRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("news article not found")
		}
		return nil, err
	}

	// Update fields
	if req.Title != nil {
		news.Title = *req.Title
		news.Slug = utils.GenerateSlug(*req.Title)
	}
	if req.Excerpt != nil {
		news.Excerpt = *req.Excerpt
	}
	if req.Content != nil {
		news.Content = *req.Content
	}
	if req.Image != nil {
		news.Image = req.Image
	}
	if req.CategoryID != nil {
		news.CategoryID = req.CategoryID
	}
	if req.Status != nil {
		news.Status = *req.Status
		// Set published_at if changing to published
		if *req.Status == "published" && news.PublishedAt == nil {
			now := time.Now()
			news.PublishedAt = &now
		}
	}
	if req.PublishedAt != nil {
		news.PublishedAt = req.PublishedAt
	}
	if req.IsFeatured != nil {
		news.IsFeatured = *req.IsFeatured
	}
	if req.MetaTitle != nil {
		news.MetaTitle = req.MetaTitle
	}
	if req.MetaDescription != nil {
		news.MetaDescription = req.MetaDescription
	}
	if req.MetaKeywords != nil {
		news.MetaKeywords = req.MetaKeywords
	}

	if err := s.newsRepo.Update(news); err != nil {
		return nil, err
	}

	// Update tags if provided
	if req.TagIDs != nil {
		if err := s.newsRepo.UpdateTags(news.ID, req.TagIDs); err != nil {
			return nil, err
		}
	}

	// Reload with relationships
	news, err = s.newsRepo.FindByID(news.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(news)
	return &response, nil
}

func (s *newsService) Delete(id uint) error {
	_, err := s.newsRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("news article not found")
		}
		return err
	}

	return s.newsRepo.Delete(id)
}

func (s *newsService) IncrementViews(id uint) error {
	return s.newsRepo.IncrementViews(id)
}

func (s *newsService) GetFeatured(limit int) ([]models.NewsArticleResponse, error) {
	news, err := s.newsRepo.FindFeatured(limit)
	if err != nil {
		return nil, err
	}

	responses := make([]models.NewsArticleResponse, len(news))
	for i, n := range news {
		responses[i] = s.toResponse(&n)
	}

	return responses, nil
}

func (s *newsService) GetByCategory(categoryID uint, page, limit int) ([]models.NewsArticleResponse, int64, error) {
	news, total, err := s.newsRepo.FindByCategory(categoryID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.NewsArticleResponse, len(news))
	for i, n := range news {
		responses[i] = s.toResponse(&n)
	}

	return responses, total, nil
}

func (s *newsService) GetByTag(tagID uint, page, limit int) ([]models.NewsArticleResponse, int64, error) {
	news, total, err := s.newsRepo.FindByTag(tagID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.NewsArticleResponse, len(news))
	for i, n := range news {
		responses[i] = s.toResponse(&n)
	}

	return responses, total, nil
}

func (s *newsService) toResponse(news *models.NewsArticle) models.NewsArticleResponse {
	return models.NewsArticleResponse{
		ID:              news.ID,
		Title:           news.Title,
		Slug:            news.Slug,
		Excerpt:         news.Excerpt,
		Content:         news.Content,
		Image:           news.Image,
		CategoryID:      news.CategoryID,
		Category:        news.Category,
		AuthorID:        news.AuthorID,
		Author:          news.Author,
		Status:          news.Status,
		PublishedAt:     news.PublishedAt,
		Views:           news.Views,
		IsFeatured:      news.IsFeatured,
		MetaTitle:       news.MetaTitle,
		MetaDescription: news.MetaDescription,
		MetaKeywords:    news.MetaKeywords,
		Tags:            news.Tags,
		CreatedAt:       news.CreatedAt,
		UpdatedAt:       news.UpdatedAt,
	}
}
