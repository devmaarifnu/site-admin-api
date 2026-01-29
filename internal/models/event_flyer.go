package models

import "time"

type EventFlyer struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Title            string     `gorm:"type:varchar(500);not null" json:"title"`
	Description      *string    `gorm:"type:text" json:"description"`
	Image            string     `gorm:"type:varchar(500);not null" json:"image"`
	EventDate        *time.Time `gorm:"type:date;index" json:"event_date"`
	EventLocation    *string    `gorm:"type:varchar(500)" json:"event_location"`
	RegistrationURL  *string    `gorm:"type:varchar(500)" json:"registration_url"`
	ContactPerson    *string    `gorm:"type:varchar(255)" json:"contact_person"`
	ContactPhone     *string    `gorm:"type:varchar(20)" json:"contact_phone"`
	ContactEmail     *string    `gorm:"type:varchar(255)" json:"contact_email"`
	OrderNumber      int        `gorm:"default:0;index" json:"order_number"`
	IsActive         bool       `gorm:"default:true;index" json:"is_active"`
	StartDisplayDate *time.Time `gorm:"index" json:"start_display_date"`
	EndDisplayDate   *time.Time `gorm:"index" json:"end_display_date"`
	CreatedBy        *uint      `gorm:"index" json:"created_by"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// Relationships
	Creator *User `gorm:"foreignKey:CreatedBy;constraint:OnDelete:SET NULL" json:"creator,omitempty"`
}

func (EventFlyer) TableName() string {
	return "event_flayers"
}

// DTOs
type EventFlyerCreateRequest struct {
	Title            string     `json:"title" binding:"required,max=500"`
	Description      *string    `json:"description"`
	Image            string     `json:"image" binding:"required"`
	EventDate        *time.Time `json:"event_date"`
	EventLocation    *string    `json:"event_location"`
	RegistrationURL  *string    `json:"registration_url"`
	ContactPerson    *string    `json:"contact_person"`
	ContactPhone     *string    `json:"contact_phone"`
	ContactEmail     *string    `json:"contact_email"`
	OrderNumber      int        `json:"order_number"`
	IsActive         bool       `json:"is_active"`
	StartDisplayDate *time.Time `json:"start_display_date"`
	EndDisplayDate   *time.Time `json:"end_display_date"`
}

type EventFlyerUpdateRequest struct {
	Title            *string    `json:"title" binding:"omitempty,max=500"`
	Description      *string    `json:"description"`
	Image            *string    `json:"image"`
	EventDate        *time.Time `json:"event_date"`
	EventLocation    *string    `json:"event_location"`
	RegistrationURL  *string    `json:"registration_url"`
	ContactPerson    *string    `json:"contact_person"`
	ContactPhone     *string    `json:"contact_phone"`
	ContactEmail     *string    `json:"contact_email"`
	OrderNumber      *int       `json:"order_number"`
	IsActive         *bool      `json:"is_active"`
	StartDisplayDate *time.Time `json:"start_display_date"`
	EndDisplayDate   *time.Time `json:"end_display_date"`
}

type EventFlyerResponse struct {
	ID               uint       `json:"id"`
	Title            string     `json:"title"`
	Description      *string    `json:"description"`
	Image            string     `json:"image"`
	EventDate        *time.Time `json:"event_date"`
	EventLocation    *string    `json:"event_location"`
	RegistrationURL  *string    `json:"registration_url"`
	ContactPerson    *string    `json:"contact_person"`
	ContactPhone     *string    `json:"contact_phone"`
	ContactEmail     *string    `json:"contact_email"`
	OrderNumber      int        `json:"order_number"`
	IsActive         bool       `json:"is_active"`
	StartDisplayDate *time.Time `json:"start_display_date"`
	EndDisplayDate   *time.Time `json:"end_display_date"`
	CreatedBy        *uint      `json:"created_by"`
	Creator          *User      `json:"creator,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
