package services

import (
	"errors"
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"

	"gorm.io/gorm"
)

type DocumentService interface {
	GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.DocumentResponse, int64, error)
	GetByID(id uint) (*models.DocumentResponse, error)
	Create(req *models.DocumentCreateRequest, uploadedBy uint) (*models.DocumentResponse, error)
	Update(id uint, req *models.DocumentUpdateRequest) (*models.DocumentResponse, error)
	Delete(id uint) error
	IncrementDownloadCount(id uint) error
	IncrementDownloads(id uint) error
	GetStats() (map[string]interface{}, error)
	GetByCategory(categoryID uint, page, limit int) ([]models.DocumentResponse, int64, error)
	GetPublic(page, limit int) ([]models.DocumentResponse, int64, error)
}

type documentService struct {
	docRepo repositories.DocumentRepository
	cfg     *config.Config
}

func NewDocumentService(docRepo repositories.DocumentRepository, cfg *config.Config) DocumentService {
	return &documentService{docRepo: docRepo, cfg: cfg}
}

func (s *documentService) GetAll(page, limit int, search string, filters map[string]interface{}) ([]models.DocumentResponse, int64, error) {
	docs, total, err := s.docRepo.FindAll(page, limit, search, filters)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.DocumentResponse, len(docs))
	for i, d := range docs {
		responses[i] = s.toResponse(&d)
	}

	return responses, total, nil
}

func (s *documentService) GetByID(id uint) (*models.DocumentResponse, error) {
	doc, err := s.docRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("document not found")
		}
		return nil, err
	}

	response := s.toResponse(doc)
	return &response, nil
}

func (s *documentService) Create(req *models.DocumentCreateRequest, uploadedBy uint) (*models.DocumentResponse, error) {
	doc := &models.Document{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		FileName:    req.FileName,
		FilePath:    req.FileURL,
		FileType:    req.FileType,
		FileSize:    req.FileSize,
		MimeType:    req.MimeType,
		IsPublic:    req.IsPublic,
		UploadedBy:  &uploadedBy,
		Status:      req.Status,
	}

	if err := s.docRepo.Create(doc); err != nil {
		return nil, err
	}

	doc, err := s.docRepo.FindByID(doc.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(doc)
	return &response, nil
}

func (s *documentService) Update(id uint, req *models.DocumentUpdateRequest) (*models.DocumentResponse, error) {
	doc, err := s.docRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("document not found")
		}
		return nil, err
	}

	if req.Title != nil {
		doc.Title = *req.Title
	}
	if req.Description != nil {
		doc.Description = req.Description
	}
	if req.CategoryID != nil {
		doc.CategoryID = req.CategoryID
	}
	if req.IsPublic != nil {
		doc.IsPublic = *req.IsPublic
	}
	if req.Status != nil {
		doc.Status = *req.Status
	}

	if err := s.docRepo.Update(doc); err != nil {
		return nil, err
	}

	doc, err = s.docRepo.FindByID(doc.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(doc)
	return &response, nil
}

func (s *documentService) Delete(id uint) error {
	_, err := s.docRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("document not found")
		}
		return err
	}

	return s.docRepo.Delete(id)
}

func (s *documentService) IncrementDownloadCount(id uint) error {
	return s.docRepo.IncrementDownloadCount(id)
}

func (s *documentService) IncrementDownloads(id uint) error {
	return s.docRepo.IncrementDownloadCount(id)
}

func (s *documentService) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Get total documents
	docs, _, err := s.docRepo.FindAll(1, 1, "", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	stats["total_documents"] = len(docs)
	stats["total_downloads"] = 0 // Placeholder, implement actual logic if needed

	return stats, nil
}

func (s *documentService) GetByCategory(categoryID uint, page, limit int) ([]models.DocumentResponse, int64, error) {
	docs, total, err := s.docRepo.FindByCategory(categoryID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.DocumentResponse, len(docs))
	for i, d := range docs {
		responses[i] = s.toResponse(&d)
	}

	return responses, total, nil
}

func (s *documentService) GetPublic(page, limit int) ([]models.DocumentResponse, int64, error) {
	docs, total, err := s.docRepo.FindPublic(page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]models.DocumentResponse, len(docs))
	for i, d := range docs {
		responses[i] = s.toResponse(&d)
	}

	return responses, total, nil
}

func (s *documentService) toResponse(doc *models.Document) models.DocumentResponse {
	return models.DocumentResponse{
		ID:            doc.ID,
		Title:         doc.Title,
		Description:   doc.Description,
		CategoryID:    doc.CategoryID,
		Category:      doc.Category,
		FileName:      doc.FileName,
		FileURL:       doc.FilePath,
		FileType:      doc.FileType,
		FileSize:      doc.FileSize,
		MimeType:      doc.MimeType,
		DownloadCount: doc.DownloadCount,
		IsPublic:      doc.IsPublic,
		UploadedBy:    doc.UploadedBy,
		Uploader:      doc.Uploader,
		Status:        doc.Status,
		CreatedAt:     doc.CreatedAt,
		UpdatedAt:     doc.UpdatedAt,
	}
}
