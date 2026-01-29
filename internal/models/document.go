package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"type:varchar(500);not null" json:"title"`
	Description   *string        `gorm:"type:text" json:"description"`
	CategoryID    *uint          `gorm:"index" json:"category_id"`
	FileName      string         `gorm:"type:varchar(255);not null" json:"file_name"`
	FilePath      string         `gorm:"type:varchar(500);not null" json:"file_path"`
	FileType      string         `gorm:"type:varchar(50);not null;index" json:"file_type"`
	FileSize      uint64         `gorm:"not null" json:"file_size"` // in bytes
	MimeType      *string        `gorm:"type:varchar(100)" json:"mime_type"`
	DownloadCount uint           `gorm:"default:0;index" json:"download_count"`
	IsPublic      bool           `gorm:"default:true;index" json:"is_public"`
	UploadedBy    *uint          `gorm:"index" json:"uploaded_by"`
	Status        string         `gorm:"type:enum('active','archived');default:'active';index" json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relationships
	Category *Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL" json:"category,omitempty"`
	Uploader *User     `gorm:"foreignKey:UploadedBy;constraint:OnDelete:SET NULL" json:"uploader,omitempty"`
}

func (Document) TableName() string {
	return "documents"
}

// DTOs
type DocumentCreateRequest struct {
	Title       string  `json:"title" binding:"required,max=500"`
	Description *string `json:"description"`
	CategoryID  *uint   `json:"category_id"`
	FileName    string  `json:"file_name" binding:"required"`
	FilePath    string  `json:"file_path" binding:"required"`
	FileType    string  `json:"file_type" binding:"required"`
	FileSize    uint64  `json:"file_size" binding:"required"`
	MimeType    *string `json:"mime_type"`
	IsPublic    bool    `json:"is_public"`
	Status      string  `json:"status" binding:"required,oneof=active archived"`
}

type DocumentUpdateRequest struct {
	Title       *string `json:"title" binding:"omitempty,max=500"`
	Description *string `json:"description"`
	CategoryID  *uint   `json:"category_id"`
	FileURL     *string `json:"file_url"`
	IsPublic    *bool   `json:"is_public"`
	Status      *string `json:"status" binding:"omitempty,oneof=active archived"`
}

type DocumentResponse struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Description   *string   `json:"description"`
	CategoryID    *uint     `json:"category_id"`
	Category      *Category `json:"category,omitempty"`
	FileName      string    `json:"file_name"`
	FilePath      string    `json:"file_path"`
	FileType      string    `json:"file_type"`
	FileSize      uint64    `json:"file_size"`
	MimeType      *string   `json:"mime_type"`
	DownloadCount uint      `json:"download_count"`
	IsPublic      bool      `json:"is_public"`
	UploadedBy    *uint     `json:"uploaded_by"`
	Uploader      *User     `json:"uploader,omitempty"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
