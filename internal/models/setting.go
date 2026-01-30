package models

import "time"

type Setting struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SettingKey   string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"setting_key"`
	SettingValue *string   `gorm:"type:text" json:"setting_value"`
	SettingType  string    `gorm:"type:enum('string','text','number','boolean','json');default:'string'" json:"setting_type"`
	SettingGroup string    `gorm:"type:varchar(50);default:'general';index" json:"setting_group"`
	Description  *string   `gorm:"type:text" json:"description"`
	IsPublic     bool      `gorm:"default:false;index" json:"is_public"` // Can be accessed without auth
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Setting) TableName() string {
	return "settings"
}

// DTOs
type SettingCreateRequest struct {
	SettingKey   string  `json:"setting_key" binding:"required,max=100"`
	SettingValue *string `json:"setting_value"`
	SettingType  string  `json:"setting_type" binding:"required,oneof=string text number boolean json"`
	SettingGroup string  `json:"setting_group"`
	Description  *string `json:"description"`
	IsPublic     bool    `json:"is_public"`
}

type SettingUpdateRequest struct {
	SettingKey   string  `json:"setting_key" binding:"required"`
	SettingValue *string `json:"setting_value"`
	SettingType  *string `json:"setting_type" binding:"omitempty,oneof=string text number boolean json"`
	SettingGroup *string `json:"setting_group"`
	Description  *string `json:"description"`
	IsPublic     *bool   `json:"is_public"`
}

type SettingResponse struct {
	ID           uint      `json:"id"`
	SettingKey   string    `json:"setting_key"`
	SettingValue *string   `json:"setting_value"`
	SettingType  string    `json:"setting_type"`
	SettingGroup string    `json:"setting_group"`
	Description  *string   `json:"description"`
	IsPublic     bool      `json:"is_public"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
