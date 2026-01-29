package models

import "time"

type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Tag) TableName() string {
	return "tags"
}

// DTOs
type TagCreateRequest struct {
	Name string `json:"name" binding:"required,max=100"`
}

type TagUpdateRequest struct {
	Name *string `json:"name" binding:"omitempty,max=100"`
}

type TagResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
