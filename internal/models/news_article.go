package models

import (
	"time"

	"gorm.io/gorm"
)

type NewsArticle struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"type:varchar(500);not null" json:"title"`
	Slug            string         `gorm:"type:varchar(500);uniqueIndex;not null" json:"slug"`
	Excerpt         string         `gorm:"type:text;not null" json:"excerpt"`
	Content         string         `gorm:"type:longtext;not null" json:"content"`
	Image           *string        `gorm:"type:varchar(500)" json:"image"`
	CategoryID      *uint          `gorm:"index" json:"category_id"`
	AuthorID        *uint          `gorm:"index" json:"author_id"`
	Status          string         `gorm:"type:enum('draft','published','archived');default:'draft';index" json:"status"`
	PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
	Views           uint           `gorm:"default:0;index" json:"views"`
	IsFeatured      bool           `gorm:"default:false;index" json:"is_featured"`
	MetaTitle       *string        `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription *string        `gorm:"type:text" json:"meta_description"`
	MetaKeywords    *string        `gorm:"type:varchar(500)" json:"meta_keywords"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relationships
	Category *Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL" json:"category,omitempty"`
	Author   *User     `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL" json:"author,omitempty"`
	Tags     []Tag     `gorm:"many2many:news_tags;" json:"tags,omitempty"`
}

func (NewsArticle) TableName() string {
	return "news_articles"
}

// DTOs
type NewsArticleCreateRequest struct {
	Title           string    `json:"title" binding:"required,max=500"`
	Excerpt         string    `json:"excerpt" binding:"required"`
	Content         string    `json:"content" binding:"required"`
	Image           *string   `json:"image"`
	CategoryID      *uint     `json:"category_id"`
	Status          string    `json:"status" binding:"required,oneof=draft published archived"`
	PublishedAt     *time.Time `json:"published_at"`
	IsFeatured      bool      `json:"is_featured"`
	MetaTitle       *string   `json:"meta_title"`
	MetaDescription *string   `json:"meta_description"`
	MetaKeywords    *string   `json:"meta_keywords"`
	TagIDs          []uint    `json:"tag_ids"`
}

type NewsArticleUpdateRequest struct {
	Title           *string    `json:"title" binding:"omitempty,max=500"`
	Excerpt         *string    `json:"excerpt"`
	Content         *string    `json:"content"`
	Image           *string    `json:"image"`
	CategoryID      *uint      `json:"category_id"`
	Status          *string    `json:"status" binding:"omitempty,oneof=draft published archived"`
	PublishedAt     *time.Time `json:"published_at"`
	IsFeatured      *bool      `json:"is_featured"`
	MetaTitle       *string    `json:"meta_title"`
	MetaDescription *string    `json:"meta_description"`
	MetaKeywords    *string    `json:"meta_keywords"`
	TagIDs          []uint     `json:"tag_ids"`
}

type NewsArticleResponse struct {
	ID              uint        `json:"id"`
	Title           string      `json:"title"`
	Slug            string      `json:"slug"`
	Excerpt         string      `json:"excerpt"`
	Content         string      `json:"content"`
	Image           *string     `json:"image"`
	CategoryID      *uint       `json:"category_id"`
	Category        *Category   `json:"category,omitempty"`
	AuthorID        *uint       `json:"author_id"`
	Author          *User       `json:"author,omitempty"`
	Status          string      `json:"status"`
	PublishedAt     *time.Time  `json:"published_at"`
	Views           uint        `json:"views"`
	IsFeatured      bool        `json:"is_featured"`
	MetaTitle       *string     `json:"meta_title"`
	MetaDescription *string     `json:"meta_description"`
	MetaKeywords    *string     `json:"meta_keywords"`
	Tags            []Tag       `json:"tags,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
