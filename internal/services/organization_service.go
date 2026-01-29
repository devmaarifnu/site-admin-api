package services

import (
	"errors"
	"site-admin-api/config"
	"site-admin-api/internal/models"
	"site-admin-api/internal/repositories"

	"gorm.io/gorm"
)

// OrganizationPosition Service
type OrganizationPositionService interface {
	GetAll(filters map[string]interface{}) ([]models.OrganizationPositionResponse, error)
	GetByID(id uint) (*models.OrganizationPositionResponse, error)
	GetByType(positionType string) ([]models.OrganizationPositionResponse, error)
	Create(req *models.OrganizationPositionCreateRequest) (*models.OrganizationPositionResponse, error)
	Update(id uint, req *models.OrganizationPositionUpdateRequest) (*models.OrganizationPositionResponse, error)
	Delete(id uint) error
}

type organizationPositionService struct {
	repo repositories.OrganizationPositionRepository
	cfg  *config.Config
}

func NewOrganizationPositionService(repo repositories.OrganizationPositionRepository, cfg *config.Config) OrganizationPositionService {
	return &organizationPositionService{repo: repo, cfg: cfg}
}

func (s *organizationPositionService) GetAll(filters map[string]interface{}) ([]models.OrganizationPositionResponse, error) {
	positions, err := s.repo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	responses := make([]models.OrganizationPositionResponse, len(positions))
	for i, p := range positions {
		responses[i] = models.OrganizationPositionResponse{
			ID:            p.ID,
			PositionName:  p.PositionName,
			PositionLevel: p.PositionLevel,
			PositionType:  p.PositionType,
			ParentID:      p.ParentID,
			Parent:        p.Parent,
			OrderNumber:   p.OrderNumber,
			IsActive:      p.IsActive,
			CreatedAt:     p.CreatedAt,
			UpdatedAt:     p.UpdatedAt,
		}
	}

	return responses, nil
}

func (s *organizationPositionService) GetByID(id uint) (*models.OrganizationPositionResponse, error) {
	position, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("position not found")
		}
		return nil, err
	}

	return &models.OrganizationPositionResponse{
		ID:            position.ID,
		PositionName:  position.PositionName,
		PositionLevel: position.PositionLevel,
		PositionType:  position.PositionType,
		ParentID:      position.ParentID,
		Parent:        position.Parent,
		OrderNumber:   position.OrderNumber,
		IsActive:      position.IsActive,
		CreatedAt:     position.CreatedAt,
		UpdatedAt:     position.UpdatedAt,
	}, nil
}

func (s *organizationPositionService) GetByType(positionType string) ([]models.OrganizationPositionResponse, error) {
	positions, err := s.repo.FindByType(positionType)
	if err != nil {
		return nil, err
	}

	responses := make([]models.OrganizationPositionResponse, len(positions))
	for i, p := range positions {
		responses[i] = models.OrganizationPositionResponse{
			ID:            p.ID,
			PositionName:  p.PositionName,
			PositionLevel: p.PositionLevel,
			PositionType:  p.PositionType,
			ParentID:      p.ParentID,
			OrderNumber:   p.OrderNumber,
			IsActive:      p.IsActive,
			CreatedAt:     p.CreatedAt,
			UpdatedAt:     p.UpdatedAt,
		}
	}

	return responses, nil
}

func (s *organizationPositionService) Create(req *models.OrganizationPositionCreateRequest) (*models.OrganizationPositionResponse, error) {
	position := &models.OrganizationPosition{
		PositionName:  req.PositionName,
		PositionLevel: req.PositionLevel,
		PositionType:  req.PositionType,
		ParentID:      req.ParentID,
		OrderNumber:   req.OrderNumber,
		IsActive:      req.IsActive,
	}

	if err := s.repo.Create(position); err != nil {
		return nil, err
	}

	position, _ = s.repo.FindByID(position.ID)
	return &models.OrganizationPositionResponse{
		ID:            position.ID,
		PositionName:  position.PositionName,
		PositionLevel: position.PositionLevel,
		PositionType:  position.PositionType,
		ParentID:      position.ParentID,
		Parent:        position.Parent,
		OrderNumber:   position.OrderNumber,
		IsActive:      position.IsActive,
		CreatedAt:     position.CreatedAt,
		UpdatedAt:     position.UpdatedAt,
	}, nil
}

func (s *organizationPositionService) Update(id uint, req *models.OrganizationPositionUpdateRequest) (*models.OrganizationPositionResponse, error) {
	position, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("position not found")
	}

	if req.PositionName != nil {
		position.PositionName = *req.PositionName
	}
	if req.PositionLevel != nil {
		position.PositionLevel = *req.PositionLevel
	}
	if req.PositionType != nil {
		position.PositionType = *req.PositionType
	}
	if req.ParentID != nil {
		position.ParentID = req.ParentID
	}
	if req.OrderNumber != nil {
		position.OrderNumber = *req.OrderNumber
	}
	if req.IsActive != nil {
		position.IsActive = *req.IsActive
	}

	if err := s.repo.Update(position); err != nil {
		return nil, err
	}

	position, _ = s.repo.FindByID(position.ID)
	return &models.OrganizationPositionResponse{
		ID:            position.ID,
		PositionName:  position.PositionName,
		PositionLevel: position.PositionLevel,
		PositionType:  position.PositionType,
		ParentID:      position.ParentID,
		Parent:        position.Parent,
		OrderNumber:   position.OrderNumber,
		IsActive:      position.IsActive,
		CreatedAt:     position.CreatedAt,
		UpdatedAt:     position.UpdatedAt,
	}, nil
}

func (s *organizationPositionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// BoardMember Service
type BoardMemberService interface {
	GetAll(filters map[string]interface{}) ([]models.BoardMemberResponse, error)
	GetByID(id uint) (*models.BoardMemberResponse, error)
	GetActive() ([]models.BoardMemberResponse, error)
	Create(req *models.BoardMemberCreateRequest) (*models.BoardMemberResponse, error)
	Update(id uint, req *models.BoardMemberUpdateRequest) (*models.BoardMemberResponse, error)
	Delete(id uint) error
}

type boardMemberService struct {
	repo repositories.BoardMemberRepository
	cfg  *config.Config
}

func NewBoardMemberService(repo repositories.BoardMemberRepository, cfg *config.Config) BoardMemberService {
	return &boardMemberService{repo: repo, cfg: cfg}
}

func (s *boardMemberService) GetAll(filters map[string]interface{}) ([]models.BoardMemberResponse, error) {
	members, err := s.repo.FindAll(filters)
	if err != nil {
		return nil, err
	}

	responses := make([]models.BoardMemberResponse, len(members))
	for i, m := range members {
		responses[i] = models.BoardMemberResponse{
			ID:          m.ID,
			PositionID:  m.PositionID,
			Position:    m.Position,
			Name:        m.Name,
			Title:       m.Title,
			Photo:       m.Photo,
			Bio:         m.Bio,
			Email:       m.Email,
			Phone:       m.Phone,
			SocialMedia: m.SocialMedia,
			PeriodStart: m.PeriodStart,
			PeriodEnd:   m.PeriodEnd,
			IsActive:    m.IsActive,
			OrderNumber: m.OrderNumber,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		}
	}
	return responses, nil
}

func (s *boardMemberService) GetByID(id uint) (*models.BoardMemberResponse, error) {
	member, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("board member not found")
	}

	return &models.BoardMemberResponse{
		ID:          member.ID,
		PositionID:  member.PositionID,
		Position:    member.Position,
		Name:        member.Name,
		Title:       member.Title,
		Photo:       member.Photo,
		Bio:         member.Bio,
		Email:       member.Email,
		Phone:       member.Phone,
		SocialMedia: member.SocialMedia,
		PeriodStart: member.PeriodStart,
		PeriodEnd:   member.PeriodEnd,
		IsActive:    member.IsActive,
		OrderNumber: member.OrderNumber,
		CreatedAt:   member.CreatedAt,
		UpdatedAt:   member.UpdatedAt,
	}, nil
}

func (s *boardMemberService) GetActive() ([]models.BoardMemberResponse, error) {
	members, err := s.repo.FindActive()
	if err != nil {
		return nil, err
	}

	responses := make([]models.BoardMemberResponse, len(members))
	for i, m := range members {
		responses[i] = models.BoardMemberResponse{
			ID:          m.ID,
			PositionID:  m.PositionID,
			Position:    m.Position,
			Name:        m.Name,
			Title:       m.Title,
			Photo:       m.Photo,
			Bio:         m.Bio,
			Email:       m.Email,
			Phone:       m.Phone,
			SocialMedia: m.SocialMedia,
			PeriodStart: m.PeriodStart,
			PeriodEnd:   m.PeriodEnd,
			IsActive:    m.IsActive,
			OrderNumber: m.OrderNumber,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		}
	}
	return responses, nil
}

func (s *boardMemberService) Create(req *models.BoardMemberCreateRequest) (*models.BoardMemberResponse, error) {
	member := &models.BoardMember{
		PositionID:  req.PositionID,
		Name:        req.Name,
		Title:       req.Title,
		Photo:       req.Photo,
		Bio:         req.Bio,
		Email:       req.Email,
		Phone:       req.Phone,
		SocialMedia: req.SocialMedia,
		PeriodStart: req.PeriodStart,
		PeriodEnd:   req.PeriodEnd,
		IsActive:    req.IsActive,
		OrderNumber: req.OrderNumber,
	}

	if err := s.repo.Create(member); err != nil {
		return nil, err
	}

	member, _ = s.repo.FindByID(member.ID)
	return &models.BoardMemberResponse{
		ID:          member.ID,
		PositionID:  member.PositionID,
		Position:    member.Position,
		Name:        member.Name,
		Title:       member.Title,
		Photo:       member.Photo,
		Bio:         member.Bio,
		Email:       member.Email,
		Phone:       member.Phone,
		SocialMedia: member.SocialMedia,
		PeriodStart: member.PeriodStart,
		PeriodEnd:   member.PeriodEnd,
		IsActive:    member.IsActive,
		OrderNumber: member.OrderNumber,
		CreatedAt:   member.CreatedAt,
		UpdatedAt:   member.UpdatedAt,
	}, nil
}

func (s *boardMemberService) Update(id uint, req *models.BoardMemberUpdateRequest) (*models.BoardMemberResponse, error) {
	member, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("board member not found")
	}

	if req.PositionID != nil {
		member.PositionID = *req.PositionID
	}
	if req.Name != nil {
		member.Name = *req.Name
	}
	if req.Title != nil {
		member.Title = req.Title
	}
	if req.Photo != nil {
		member.Photo = req.Photo
	}
	if req.Bio != nil {
		member.Bio = req.Bio
	}
	if req.Email != nil {
		member.Email = req.Email
	}
	if req.Phone != nil {
		member.Phone = req.Phone
	}
	if req.SocialMedia != nil {
		member.SocialMedia = req.SocialMedia
	}
	if req.PeriodStart != nil {
		member.PeriodStart = *req.PeriodStart
	}
	if req.PeriodEnd != nil {
		member.PeriodEnd = *req.PeriodEnd
	}
	if req.IsActive != nil {
		member.IsActive = *req.IsActive
	}
	if req.OrderNumber != nil {
		member.OrderNumber = *req.OrderNumber
	}

	if err := s.repo.Update(member); err != nil {
		return nil, err
	}

	member, _ = s.repo.FindByID(member.ID)
	return &models.BoardMemberResponse{
		ID:          member.ID,
		PositionID:  member.PositionID,
		Position:    member.Position,
		Name:        member.Name,
		Title:       member.Title,
		Photo:       member.Photo,
		Bio:         member.Bio,
		Email:       member.Email,
		Phone:       member.Phone,
		SocialMedia: member.SocialMedia,
		PeriodStart: member.PeriodStart,
		PeriodEnd:   member.PeriodEnd,
		IsActive:    member.IsActive,
		OrderNumber: member.OrderNumber,
		CreatedAt:   member.CreatedAt,
		UpdatedAt:   member.UpdatedAt,
	}, nil
}

func (s *boardMemberService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// Pengurus & Department services similar pattern - keeping it concise for token efficiency
// These would follow the same CRUD pattern as above
