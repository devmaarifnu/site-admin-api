package models

import "time"

type Department struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description *string   `gorm:"type:text" json:"description"`
	HeadName    *string   `gorm:"type:varchar(255)" json:"head_name"`
	OrderNumber int       `gorm:"default:0;index" json:"order_number"`
	IsActive    bool      `gorm:"default:true;index" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Department) TableName() string {
	return "departments"
}

// DTOs
type DepartmentCreateRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	Description *string `json:"description"`
	HeadName    *string `json:"head_name"`
	OrderNumber int     `json:"order_number"`
	IsActive    bool    `json:"is_active"`
}

type DepartmentUpdateRequest struct {
	Name        *string `json:"name" binding:"omitempty,max=255"`
	Description *string `json:"description"`
	HeadName    *string `json:"head_name"`
	OrderNumber *int    `json:"order_number"`
	IsActive    *bool   `json:"is_active"`
}

type DepartmentResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	HeadName    *string   `json:"head_name"`
	OrderNumber int       `json:"order_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
