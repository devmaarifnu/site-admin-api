package models

import "time"

type OrganizationPosition struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PositionName  string    `gorm:"type:varchar(255);not null" json:"position_name"`
	PositionLevel int       `gorm:"not null;index" json:"position_level"` // 1=Ketua, 2=Wakil, 3=Sekretaris, 4=Bendahara, 5=Bidang
	PositionType  string    `gorm:"type:enum('ketua','wakil','sekretaris','bendahara','bidang');not null;index" json:"position_type"`
	ParentID      *uint     `gorm:"index" json:"parent_id"`
	OrderNumber   int       `gorm:"default:0;index" json:"order_number"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relationships
	Parent *OrganizationPosition `gorm:"foreignKey:ParentID;constraint:OnDelete:SET NULL" json:"parent,omitempty"`
}

func (OrganizationPosition) TableName() string {
	return "organization_positions"
}

// DTOs
type OrganizationPositionCreateRequest struct {
	PositionName  string  `json:"position_name" binding:"required,max=255"`
	PositionLevel int     `json:"position_level" binding:"required"`
	PositionType  string  `json:"position_type" binding:"required,oneof=ketua wakil sekretaris bendahara bidang"`
	ParentID      *uint   `json:"parent_id"`
	OrderNumber   int     `json:"order_number"`
	IsActive      bool    `json:"is_active"`
}

type OrganizationPositionUpdateRequest struct {
	PositionName  *string `json:"position_name" binding:"omitempty,max=255"`
	PositionLevel *int    `json:"position_level"`
	PositionType  *string `json:"position_type" binding:"omitempty,oneof=ketua wakil sekretaris bendahara bidang"`
	ParentID      *uint   `json:"parent_id"`
	OrderNumber   *int    `json:"order_number"`
	IsActive      *bool   `json:"is_active"`
}

type OrganizationPositionResponse struct {
	ID            uint                      `json:"id"`
	PositionName  string                    `json:"position_name"`
	PositionLevel int                       `json:"position_level"`
	PositionType  string                    `json:"position_type"`
	ParentID      *uint                     `json:"parent_id"`
	Parent        *OrganizationPosition     `json:"parent,omitempty"`
	OrderNumber   int                       `json:"order_number"`
	IsActive      bool                      `json:"is_active"`
	CreatedAt     time.Time                 `json:"created_at"`
	UpdatedAt     time.Time                 `json:"updated_at"`
}
