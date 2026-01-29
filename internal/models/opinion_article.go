package models

import (
	"time"

	"gorm.io/gorm"
)

type OpinionArticle struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"type:varchar(500);not null" json:"title"`
	Slug            string         `gorm:"type:varchar(500);uniqueIndex;not null" json:"slug"`
	Excerpt         string         `gorm:"type:text;not null" json:"excerpt"`
	Content         string         `gorm:"type:longtext;not null" json:"content"`
	Image           *string        `gorm:"type:varchar(500)" json:"image"`
	AuthorName      string         `gorm:"type:varchar(255);not null" json:"author_name"`
	AuthorTitle     *string        `gorm:"type:varchar(255)" json:"author_title"`
	AuthorImage     *string        `gorm:"type:varchar(500)" json:"author_image"`
	AuthorBio       *string        `gorm:"type:text" json:"author_bio"`
	Status          string         `gorm:"type:enum('draft','published','archived');default:'draft';index" json:"status"`
	PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
	Views           uint           `gorm:"default:0;index" json:"views"`
	IsFeatured      bool           `gorm:"default:false;index" json:"is_featured"`
	MetaTitle       *string        `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription *string        `gorm:"type:text" json:"meta_description"`
	MetaKeywords    *string        `gorm:"type:varchar(500)" json:"meta_keywords"`
	CreatedBy       *uint          `gorm:"index" json:"created_by"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relationships
	Creator *User `gorm:"foreignKey:CreatedBy;constraint:OnDelete:SET NULL" json:"creator,omitempty"`
	Tags    []Tag `gorm:"many2many:opinion_tags;" json:"tags,omitempty"`
}

func (OpinionArticle) TableName() string {
	return "opinion_articles"
}

// DTOs
type OpinionArticleCreateRequest struct {
	Title           string     `json:"title" binding:"required,max=500"`
	Excerpt         string     `json:"excerpt" binding:"required"`
	Content         string     `json:"content" binding:"required"`
	Image           *string    `json:"image"`
	AuthorName      string     `json:"author_name" binding:"required,max=255"`
	AuthorTitle     *string    `json:"author_title"`
	AuthorImage     *string    `json:"author_image"`
	AuthorBio       *string    `json:"author_bio"`
	Status          string     `json:"status" binding:"required,oneof=draft published archived"`
	PublishedAt     *time.Time `json:"published_at"`
	IsFeatured      bool       `json:"is_featured"`
	MetaTitle       *string    `json:"meta_title"`
	MetaDescription *string    `json:"meta_description"`
	MetaKeywords    *string    `json:"meta_keywords"`
	TagIDs          []uint     `json:"tag_ids"`
}

type OpinionArticleUpdateRequest struct {
	Title           *string    `json:"title" binding:"omitempty,max=500"`
	Excerpt         *string    `json:"excerpt"`
	Content         *string    `json:"content"`
	Image           *string    `json:"image"`
	AuthorName      *string    `json:"author_name" binding:"omitempty,max=255"`
	AuthorTitle     *string    `json:"author_title"`
	AuthorImage     *string    `json:"author_image"`
	AuthorBio       *string    `json:"author_bio"`
	Status          *string    `json:"status" binding:"omitempty,oneof=draft published archived"`
	PublishedAt     *time.Time `json:"published_at"`
	IsFeatured      *bool      `json:"is_featured"`
	MetaTitle       *string    `json:"meta_title"`
	MetaDescription *string    `json:"meta_description"`
	MetaKeywords    *string    `json:"meta_keywords"`
	TagIDs          []uint     `json:"tag_ids"`
}

type OpinionArticleResponse struct {
	ID              uint       `json:"id"`
	Title           string     `json:"title"`
	Slug            string     `json:"slug"`
	Excerpt         string     `json:"excerpt"`
	Content         string     `json:"content"`
	Image           *string    `json:"image"`
	AuthorName      string     `json:"author_name"`
	AuthorTitle     *string    `json:"author_title"`
	AuthorImage     *string    `json:"author_image"`
	AuthorBio       *string    `json:"author_bio"`
	Status          string     `json:"status"`
	PublishedAt     *time.Time `json:"published_at"`
	Views           uint       `json:"views"`
	IsFeatured      bool       `json:"is_featured"`
	MetaTitle       *string    `json:"meta_title"`
	MetaDescription *string    `json:"meta_description"`
	MetaKeywords    *string    `json:"meta_keywords"`
	CreatedBy       *uint      `json:"created_by"`
	Creator         *User      `json:"creator,omitempty"`
	Tags            []Tag      `json:"tags,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
