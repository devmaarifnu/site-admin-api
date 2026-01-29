package models

import "time"

type Page struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Slug            string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Title           string    `gorm:"type:varchar(500);not null" json:"title"`
	Content         *string   `gorm:"type:longtext" json:"content"`
	Metadata        JSONB     `gorm:"type:json" json:"metadata"` // Flexible content based on page type
	Template        string    `gorm:"type:varchar(100);default:'default'" json:"template"`
	IsActive        bool      `gorm:"default:true;index" json:"is_active"`
	MetaTitle       *string   `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription *string   `gorm:"type:text" json:"meta_description"`
	MetaKeywords    *string   `gorm:"type:varchar(500)" json:"meta_keywords"`
	LastUpdatedBy   *uint     `gorm:"index" json:"last_updated_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// Relationships
	LastEditor *User `gorm:"foreignKey:LastUpdatedBy;constraint:OnDelete:SET NULL" json:"last_editor,omitempty"`
}

func (Page) TableName() string {
	return "pages"
}

// DTOs
type PageCreateRequest struct {
	Slug            string  `json:"slug" binding:"required,max=255"`
	Title           string  `json:"title" binding:"required,max=500"`
	Content         *string `json:"content"`
	Metadata        JSONB   `json:"metadata"`
	Template        string  `json:"template"`
	IsActive        bool    `json:"is_active"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaKeywords    *string `json:"meta_keywords"`
}

type PageUpdateRequest struct {
	Slug            *string `json:"slug" binding:"omitempty,max=255"`
	Title           *string `json:"title" binding:"omitempty,max=500"`
	Content         *string `json:"content"`
	Metadata        JSONB   `json:"metadata"`
	Template        *string `json:"template"`
	IsActive        *bool   `json:"is_active"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaKeywords    *string `json:"meta_keywords"`
}

type PageResponse struct {
	ID              uint      `json:"id"`
	Slug            string    `json:"slug"`
	Title           string    `json:"title"`
	Content         *string   `json:"content"`
	Metadata        JSONB     `json:"metadata"`
	Template        string    `json:"template"`
	IsActive        bool      `json:"is_active"`
	MetaTitle       *string   `json:"meta_title"`
	MetaDescription *string   `json:"meta_description"`
	MetaKeywords    *string   `json:"meta_keywords"`
	LastUpdatedBy   *uint     `json:"last_updated_by"`
	LastEditor      *User     `json:"last_editor,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
