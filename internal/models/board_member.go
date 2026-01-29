package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Custom type for JSON fields
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONB)
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

type BoardMember struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PositionID  uint      `gorm:"not null;index" json:"position_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Title       *string   `gorm:"type:varchar(255)" json:"title"` // Gelar akademik
	Photo       *string   `gorm:"type:varchar(500)" json:"photo"`
	Bio         *string   `gorm:"type:text" json:"bio"`
	Email       *string   `gorm:"type:varchar(255)" json:"email"`
	Phone       *string   `gorm:"type:varchar(20)" json:"phone"`
	SocialMedia JSONB     `gorm:"type:json" json:"social_media"`
	PeriodStart int       `gorm:"type:year;not null;index" json:"period_start"`
	PeriodEnd   int       `gorm:"type:year;not null;index" json:"period_end"`
	IsActive    bool      `gorm:"default:true;index" json:"is_active"`
	OrderNumber int       `gorm:"default:0;index" json:"order_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Position *OrganizationPosition `gorm:"foreignKey:PositionID;constraint:OnDelete:CASCADE" json:"position,omitempty"`
}

func (BoardMember) TableName() string {
	return "board_members"
}

// DTOs
type BoardMemberCreateRequest struct {
	PositionID  uint    `json:"position_id" binding:"required"`
	Name        string  `json:"name" binding:"required,max=255"`
	Title       *string `json:"title"`
	Photo       *string `json:"photo"`
	Bio         *string `json:"bio"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	SocialMedia JSONB   `json:"social_media"`
	PeriodStart int     `json:"period_start" binding:"required"`
	PeriodEnd   int     `json:"period_end" binding:"required"`
	IsActive    bool    `json:"is_active"`
	OrderNumber int     `json:"order_number"`
}

type BoardMemberUpdateRequest struct {
	PositionID  *uint   `json:"position_id"`
	Name        *string `json:"name" binding:"omitempty,max=255"`
	Title       *string `json:"title"`
	Photo       *string `json:"photo"`
	Bio         *string `json:"bio"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	SocialMedia JSONB   `json:"social_media"`
	PeriodStart *int    `json:"period_start"`
	PeriodEnd   *int    `json:"period_end"`
	IsActive    *bool   `json:"is_active"`
	OrderNumber *int    `json:"order_number"`
}

type BoardMemberResponse struct {
	ID          uint                  `json:"id"`
	PositionID  uint                  `json:"position_id"`
	Position    *OrganizationPosition `json:"position,omitempty"`
	Name        string                `json:"name"`
	Title       *string               `json:"title"`
	Photo       *string               `json:"photo"`
	Bio         *string               `json:"bio"`
	Email       *string               `json:"email"`
	Phone       *string               `json:"phone"`
	SocialMedia JSONB                 `json:"social_media"`
	PeriodStart int                   `json:"period_start"`
	PeriodEnd   int                   `json:"period_end"`
	IsActive    bool                  `json:"is_active"`
	OrderNumber int                   `json:"order_number"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}
