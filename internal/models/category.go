package models

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Slug        string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	Description *string   `gorm:"type:text" json:"description"`
	Type        string    `gorm:"type:enum('news','opinion','document');not null;index" json:"type"`
	Icon        *string   `gorm:"type:varchar(100)" json:"icon"`
	Color       *string   `gorm:"type:varchar(7)" json:"color"`
	IsActive    bool      `gorm:"default:true;index" json:"is_active"`
	OrderNumber int       `gorm:"default:0" json:"order_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

// DTOs
type CategoryCreateRequest struct {
	Name        string  `json:"name" binding:"required,max=100"`
	Description *string `json:"description"`
	Type        string  `json:"type" binding:"required,oneof=news opinion document"`
	Icon        *string `json:"icon"`
	Color       *string `json:"color"`
	IsActive    bool    `json:"is_active"`
	OrderNumber int     `json:"order_number"`
}

type CategoryUpdateRequest struct {
	Name        *string `json:"name" binding:"omitempty,max=100"`
	Description *string `json:"description"`
	Type        *string `json:"type" binding:"omitempty,oneof=news opinion document"`
	Icon        *string `json:"icon"`
	Color       *string `json:"color"`
	IsActive    *bool   `json:"is_active"`
	OrderNumber *int    `json:"order_number"`
}

type CategoryResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description"`
	Type        string    `json:"type"`
	Icon        *string   `json:"icon"`
	Color       *string   `json:"color"`
	IsActive    bool      `json:"is_active"`
	OrderNumber int       `json:"order_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
