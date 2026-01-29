package models

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	FileName     string         `gorm:"type:varchar(255);not null" json:"file_name"`
	OriginalName string         `gorm:"type:varchar(255);not null" json:"original_name"`
	FilePath     string         `gorm:"type:varchar(500);not null" json:"file_path"`
	FileURL      string         `gorm:"type:varchar(500);not null" json:"file_url"`
	FileType     string         `gorm:"type:varchar(50);not null;index" json:"file_type"`
	MimeType     *string        `gorm:"type:varchar(100)" json:"mime_type"`
	FileSize     uint64         `gorm:"not null" json:"file_size"` // in bytes
	Width        *uint          `json:"width"`                     // for images
	Height       *uint          `json:"height"`                    // for images
	Folder       string         `gorm:"type:varchar(100);default:'general';index" json:"folder"`
	AltText      *string        `gorm:"type:varchar(255)" json:"alt_text"`
	Caption      *string        `gorm:"type:text" json:"caption"`
	UploadedBy   *uint          `gorm:"index" json:"uploaded_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relationships
	Uploader *User `gorm:"foreignKey:UploadedBy;constraint:OnDelete:SET NULL" json:"uploader,omitempty"`
}

func (Media) TableName() string {
	return "media"
}

// DTOs
type MediaCreateRequest struct {
	FileName     string  `json:"file_name" binding:"required"`
	OriginalName string  `json:"original_name" binding:"required"`
	FilePath     string  `json:"file_path" binding:"required"`
	FileURL      string  `json:"file_url" binding:"required"`
	FileType     string  `json:"file_type" binding:"required"`
	MimeType     *string `json:"mime_type"`
	FileSize     uint64  `json:"file_size" binding:"required"`
	Width        *uint   `json:"width"`
	Height       *uint   `json:"height"`
	Folder       string  `json:"folder"`
	AltText      *string `json:"alt_text"`
	Caption      *string `json:"caption"`
}

type MediaUpdateRequest struct {
	AltText *string `json:"alt_text"`
	Caption *string `json:"caption"`
	Folder  *string `json:"folder"`
}

type MediaResponse struct {
	ID           uint      `json:"id"`
	FileName     string    `json:"file_name"`
	OriginalName string    `json:"original_name"`
	FilePath     string    `json:"file_path"`
	FileURL      string    `json:"file_url"`
	FileType     string    `json:"file_type"`
	MimeType     *string   `json:"mime_type"`
	FileSize     uint64    `json:"file_size"`
	Width        *uint     `json:"width"`
	Height       *uint     `json:"height"`
	Folder       string    `json:"folder"`
	AltText      *string   `json:"alt_text"`
	Caption      *string   `json:"caption"`
	UploadedBy   *uint     `json:"uploaded_by"`
	Uploader     *User     `json:"uploader,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
