package services

import (
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"
)

type EventFlyerService interface {
	GetAll(page, limit int, search string) ([]models.EventFlyerResponse, int64, error)
	GetByID(id uint) (*models.EventFlyerResponse, error)
	Create(req *models.EventFlyerCreateRequest, uploaderID uint) (*models.EventFlyerResponse, error)
	Update(id uint, req *models.EventFlyerUpdateRequest) (*models.EventFlyerResponse, error)
	Delete(id uint) error
}

type eventFlyerService struct {
	eventFlyerRepo repositories.EventFlyerRepository
	cfg            *config.Config
}

func NewEventFlyerService(eventFlyerRepo repositories.EventFlyerRepository, cfg *config.Config) EventFlyerService {
	return &eventFlyerService{
		eventFlyerRepo: eventFlyerRepo,
		cfg:            cfg,
	}
}

func (s *eventFlyerService) GetAll(page, limit int, search string) ([]models.EventFlyerResponse, int64, error) {
	filters := make(map[string]interface{})
	flyers, err := s.eventFlyerRepo.FindAll(filters)
	if err != nil {
		return nil, 0, err
	}

	var responses []models.EventFlyerResponse
	for _, flyer := range flyers {
		responses = append(responses, s.toResponse(&flyer))
	}

	return responses, int64(len(flyers)), nil
}

func (s *eventFlyerService) GetByID(id uint) (*models.EventFlyerResponse, error) {
	flyer, err := s.eventFlyerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(flyer)
	return &response, nil
}

func (s *eventFlyerService) Create(req *models.EventFlyerCreateRequest, uploaderID uint) (*models.EventFlyerResponse, error) {
	flyer := &models.EventFlyer{
		Title:            req.Title,
		Description:      req.Description,
		Image:            req.Image,
		EventDate:        req.EventDate,
		EventLocation:    req.EventLocation,
		RegistrationURL:  req.RegistrationURL,
		ContactPerson:    req.ContactPerson,
		ContactPhone:     req.ContactPhone,
		ContactEmail:     req.ContactEmail,
		OrderNumber:      req.OrderNumber,
		IsActive:         req.IsActive,
		StartDisplayDate: req.StartDisplayDate,
		EndDisplayDate:   req.EndDisplayDate,
		CreatedBy:        &uploaderID,
	}

	if err := s.eventFlyerRepo.Create(flyer); err != nil {
		return nil, err
	}

	// Reload with creator
	flyer, err := s.eventFlyerRepo.FindByID(flyer.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(flyer)
	return &response, nil
}

func (s *eventFlyerService) Update(id uint, req *models.EventFlyerUpdateRequest) (*models.EventFlyerResponse, error) {
	flyer, err := s.eventFlyerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.Title != nil {
		flyer.Title = *req.Title
	}
	if req.Description != nil {
		desc := *req.Description
		flyer.Description = &desc
	}
	if req.Image != nil {
		flyer.Image = *req.Image
	}
	if req.EventDate != nil {
		date := *req.EventDate
		flyer.EventDate = &date
	}
	if req.EventLocation != nil {
		loc := *req.EventLocation
		flyer.EventLocation = &loc
	}
	if req.RegistrationURL != nil {
		url := *req.RegistrationURL
		flyer.RegistrationURL = &url
	}
	if req.ContactPerson != nil {
		person := *req.ContactPerson
		flyer.ContactPerson = &person
	}
	if req.ContactPhone != nil {
		phone := *req.ContactPhone
		flyer.ContactPhone = &phone
	}
	if req.ContactEmail != nil {
		email := *req.ContactEmail
		flyer.ContactEmail = &email
	}
	if req.OrderNumber != nil {
		flyer.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		flyer.IsActive = *req.IsActive
	}
	if req.StartDisplayDate != nil {
		startDate := *req.StartDisplayDate
		flyer.StartDisplayDate = &startDate
	}
	if req.EndDisplayDate != nil {
		endDate := *req.EndDisplayDate
		flyer.EndDisplayDate = &endDate
	}

	if err := s.eventFlyerRepo.Update(flyer); err != nil {
		return nil, err
	}

	// Reload with creator
	flyer, err = s.eventFlyerRepo.FindByID(flyer.ID)
	if err != nil {
		return nil, err
	}

	response := s.toResponse(flyer)
	return &response, nil
}

func (s *eventFlyerService) Delete(id uint) error {
	return s.eventFlyerRepo.Delete(id)
}

func (s *eventFlyerService) toResponse(flyer *models.EventFlyer) models.EventFlyerResponse {
	return models.EventFlyerResponse{
		ID:               flyer.ID,
		Title:            flyer.Title,
		Description:      flyer.Description,
		Image:            flyer.Image,
		EventDate:        flyer.EventDate,
		EventLocation:    flyer.EventLocation,
		RegistrationURL:  flyer.RegistrationURL,
		ContactPerson:    flyer.ContactPerson,
		ContactPhone:     flyer.ContactPhone,
		ContactEmail:     flyer.ContactEmail,
		OrderNumber:      flyer.OrderNumber,
		IsActive:         flyer.IsActive,
		StartDisplayDate: flyer.StartDisplayDate,
		EndDisplayDate:   flyer.EndDisplayDate,
		CreatedBy:        flyer.CreatedBy,
		Creator:          flyer.Creator,
		CreatedAt:        flyer.CreatedAt,
		UpdatedAt:        flyer.UpdatedAt,
	}
}
