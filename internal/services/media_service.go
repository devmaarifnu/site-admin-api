package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type MediaService interface {
	GetAll(page, limit int, filters map[string]interface{}) ([]models.MediaResponse, int64, error)
	GetByID(id uint) (*models.MediaResponse, error)
	Create(req *models.MediaCreateRequest, uploaderID uint) (*models.MediaResponse, error)
	Delete(id uint) error
}

type mediaService struct {
	mediaRepo repositories.MediaRepository
	cfg       *config.Config
}

func NewMediaService(mediaRepo repositories.MediaRepository, cfg *config.Config) MediaService {
	return &mediaService{
		mediaRepo: mediaRepo,
		cfg:       cfg,
	}
}

func (s *mediaService) GetAll(page, limit int, filters map[string]interface{}) ([]models.MediaResponse, int64, error) {
	media, total, err := s.mediaRepo.FindAll(page, limit, filters)
	if err != nil {
		return nil, 0, err
	}

	var responses []models.MediaResponse
	for _, m := range media {
		responses = append(responses, s.toResponse(&m))
	}

	return responses, total, nil
}

func (s *mediaService) GetByID(id uint) (*models.MediaResponse, error) {
	media, err := s.mediaRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(media)
	return &response, nil
}

func (s *mediaService) Create(req *models.MediaCreateRequest, uploaderID uint) (*models.MediaResponse, error) {
	media := &models.Media{
		FileName:     req.FileName,
		OriginalName: req.OriginalName,
		FilePath:     req.FilePath,
		FileURL:      req.FileURL,
		FileType:     req.FileType,
		MimeType:     req.MimeType,
		FileSize:     req.FileSize,
		Width:        req.Width,
		Height:       req.Height,
		Folder:       req.Folder,
		AltText:      req.AltText,
		Caption:      req.Caption,
		UploadedBy:   &uploaderID,
	}

	if err := s.mediaRepo.Create(media); err != nil {
		return nil, err
	}

	// Reload with uploader
	media, err := s.mediaRepo.FindByID(media.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(media)
	return &response, nil
}

func (s *mediaService) Delete(id uint) error {
	return s.mediaRepo.Delete(id)
}

func (s *mediaService) toResponse(media *models.Media) models.MediaResponse {
	return models.MediaResponse{
		ID:           media.ID,
		FileName:     media.FileName,
		OriginalName: media.OriginalName,
		FilePath:     media.FilePath,
		FileURL:      media.FileURL,
		FileType:     media.FileType,
		MimeType:     media.MimeType,
		FileSize:     media.FileSize,
		Width:        media.Width,
		Height:       media.Height,
		Folder:       media.Folder,
		AltText:      media.AltText,
		Caption:      media.Caption,
		UploadedBy:   media.UploadedBy,
		Uploader:     media.Uploader,
		CreatedAt:    media.CreatedAt,
		UpdatedAt:    media.UpdatedAt,
	}
}
