package models

import "time"

type EditorialCouncil struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Institution string    `gorm:"type:varchar(255);not null" json:"institution"`
	Expertise   *string   `gorm:"type:varchar(500)" json:"expertise"`
	Photo       *string   `gorm:"type:varchar(500)" json:"photo"`
	Bio         *string   `gorm:"type:text" json:"bio"`
	Email       *string   `gorm:"type:varchar(255)" json:"email"`
	OrderNumber int       `gorm:"default:0;index" json:"order_number"`
	IsActive    bool      `gorm:"default:true;index" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (EditorialCouncil) TableName() string {
	return "editorial_council"
}

// DTOs
type EditorialCouncilCreateRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	Institution string  `json:"institution" binding:"required,max=255"`
	Expertise   *string `json:"expertise"`
	Photo       *string `json:"photo"`
	Bio         *string `json:"bio"`
	Email       *string `json:"email"`
	OrderNumber int     `json:"order_number"`
	IsActive    bool    `json:"is_active"`
}

type EditorialCouncilUpdateRequest struct {
	Name        *string `json:"name" binding:"omitempty,max=255"`
	Institution *string `json:"institution" binding:"omitempty,max=255"`
	Expertise   *string `json:"expertise"`
	Photo       *string `json:"photo"`
	Bio         *string `json:"bio"`
	Email       *string `json:"email"`
	OrderNumber *int    `json:"order_number"`
	IsActive    *bool   `json:"is_active"`
}

type EditorialCouncilResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Institution string    `json:"institution"`
	Expertise   *string   `json:"expertise"`
	Photo       *string   `json:"photo"`
	Bio         *string   `json:"bio"`
	Email       *string   `json:"email"`
	OrderNumber int       `json:"order_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
