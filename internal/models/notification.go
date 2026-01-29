package models

import "time"

type Notification struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null;index" json:"user_id"`
	Type      string     `gorm:"type:varchar(100);not null" json:"type"`
	Title     string     `gorm:"type:varchar(255);not null" json:"title"`
	Message   string     `gorm:"type:text;not null" json:"message"`
	Data      JSONB      `gorm:"type:json" json:"data"`
	ReadAt    *time.Time `json:"read_at"`
	CreatedAt time.Time  `gorm:"index" json:"created_at"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}

// DTOs
type NotificationCreateRequest struct {
	UserID  uint    `json:"user_id" binding:"required"`
	Type    string  `json:"type" binding:"required,max=100"`
	Title   string  `json:"title" binding:"required,max=255"`
	Message string  `json:"message" binding:"required"`
	Data    JSONB   `json:"data"`
}

type NotificationResponse struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	User      *User      `json:"user,omitempty"`
	Type      string     `json:"type"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Data      JSONB      `json:"data"`
	ReadAt    *time.Time `json:"read_at"`
	CreatedAt time.Time  `json:"created_at"`
}
