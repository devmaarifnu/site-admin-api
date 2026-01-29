package models

import "time"

type ContactMessage struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	TicketID   string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"ticket_id"`
	Name       string     `gorm:"type:varchar(255);not null" json:"name"`
	Email      string     `gorm:"type:varchar(255);not null;index" json:"email"`
	Phone      *string    `gorm:"type:varchar(20)" json:"phone"`
	Subject    string     `gorm:"type:varchar(500);not null" json:"subject"`
	Message    string     `gorm:"type:text;not null" json:"message"`
	Status     string     `gorm:"type:enum('new','read','in_progress','resolved','closed');default:'new';index" json:"status"`
	Priority   string     `gorm:"type:enum('low','medium','high','urgent');default:'medium';index" json:"priority"`
	IPAddress  *string    `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent  *string    `gorm:"type:text" json:"user_agent"`
	AssignedTo *uint      `gorm:"index" json:"assigned_to"`
	RepliedAt  *time.Time `json:"replied_at"`
	ResolvedAt *time.Time `json:"resolved_at"`
	Notes      *string    `gorm:"type:text" json:"notes"` // Internal notes
	CreatedAt  time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	// Relationships
	AssignedUser *User `gorm:"foreignKey:AssignedTo;constraint:OnDelete:SET NULL" json:"assigned_user,omitempty"`
}

func (ContactMessage) TableName() string {
	return "contact_messages"
}

// DTOs
type ContactMessageCreateRequest struct {
	Name      string  `json:"name" binding:"required,max=255"`
	Email     string  `json:"email" binding:"required,email,max=255"`
	Phone     *string `json:"phone"`
	Subject   string  `json:"subject" binding:"required,max=500"`
	Message   string  `json:"message" binding:"required"`
	IPAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
}

type ContactMessageUpdateRequest struct {
	Status     *string    `json:"status" binding:"omitempty,oneof=new read in_progress resolved closed"`
	Priority   *string    `json:"priority" binding:"omitempty,oneof=low medium high urgent"`
	AssignedTo *uint      `json:"assigned_to"`
	RepliedAt  *time.Time `json:"replied_at"`
	ResolvedAt *time.Time `json:"resolved_at"`
	Notes      *string    `json:"notes"`
}

type ContactMessageResponse struct {
	ID           uint       `json:"id"`
	TicketID     string     `json:"ticket_id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Phone        *string    `json:"phone"`
	Subject      string     `json:"subject"`
	Message      string     `json:"message"`
	Status       string     `json:"status"`
	Priority     string     `json:"priority"`
	IPAddress    *string    `json:"ip_address"`
	UserAgent    *string    `json:"user_agent"`
	AssignedTo   *uint      `json:"assigned_to"`
	AssignedUser *User      `json:"assigned_user,omitempty"`
	RepliedAt    *time.Time `json:"replied_at"`
	ResolvedAt   *time.Time `json:"resolved_at"`
	Notes        *string    `json:"notes"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
