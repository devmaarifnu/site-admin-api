package models

import "time"

type EditorialTeam struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Position    string    `gorm:"type:varchar(255);not null" json:"position"`
	RoleType    string    `gorm:"type:enum('pemimpin_redaksi','wakil_pemimpin_redaksi','redaktur_pelaksana','tim_redaksi');not null;index" json:"role_type"`
	Photo       *string   `gorm:"type:varchar(500)" json:"photo"`
	Bio         *string   `gorm:"type:text" json:"bio"`
	Email       *string   `gorm:"type:varchar(255)" json:"email"`
	Phone       *string   `gorm:"type:varchar(20)" json:"phone"`
	OrderNumber int       `gorm:"default:0;index" json:"order_number"`
	IsActive    bool      `gorm:"default:true;index" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (EditorialTeam) TableName() string {
	return "editorial_team"
}

// DTOs
type EditorialTeamCreateRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	Position    string  `json:"position" binding:"required,max=255"`
	RoleType    string  `json:"role_type" binding:"required,oneof=pemimpin_redaksi wakil_pemimpin_redaksi redaktur_pelaksana tim_redaksi"`
	Photo       *string `json:"photo"`
	Bio         *string `json:"bio"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	OrderNumber int     `json:"order_number"`
	IsActive    bool    `json:"is_active"`
}

type EditorialTeamUpdateRequest struct {
	Name        *string `json:"name" binding:"omitempty,max=255"`
	Position    *string `json:"position" binding:"omitempty,max=255"`
	RoleType    *string `json:"role_type" binding:"omitempty,oneof=pemimpin_redaksi wakil_pemimpin_redaksi redaktur_pelaksana tim_redaksi"`
	Photo       *string `json:"photo"`
	Bio         *string `json:"bio"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	OrderNumber *int    `json:"order_number"`
	IsActive    *bool   `json:"is_active"`
}

type EditorialTeamResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	RoleType    string    `json:"role_type"`
	Photo       *string   `json:"photo"`
	Bio         *string   `json:"bio"`
	Email       *string   `json:"email"`
	Phone       *string   `json:"phone"`
	OrderNumber int       `json:"order_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
