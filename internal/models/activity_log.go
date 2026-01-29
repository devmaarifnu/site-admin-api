package models

import "time"

type ActivityLog struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LogName     *string   `gorm:"type:varchar(100);index" json:"log_name"`
	Description *string   `gorm:"type:text" json:"description"`
	SubjectType *string   `gorm:"type:varchar(255);index" json:"subject_type"`
	SubjectID   *uint     `gorm:"index" json:"subject_id"`
	CauserType  *string   `gorm:"type:varchar(255);index" json:"causer_type"`
	CauserID    *uint     `gorm:"index" json:"causer_id"`
	Properties  JSONB     `gorm:"type:json" json:"properties"`
	IPAddress   *string   `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent   *string   `gorm:"type:text" json:"user_agent"`
	CreatedAt   time.Time `gorm:"index" json:"created_at"`
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}

// DTOs
type ActivityLogCreateRequest struct {
	LogName     *string `json:"log_name"`
	Description *string `json:"description"`
	SubjectType *string `json:"subject_type"`
	SubjectID   *uint   `json:"subject_id"`
	CauserType  *string `json:"causer_type"`
	CauserID    *uint   `json:"causer_id"`
	Properties  JSONB   `json:"properties"`
	IPAddress   *string `json:"ip_address"`
	UserAgent   *string `json:"user_agent"`
}

type ActivityLogResponse struct {
	ID          uint      `json:"id"`
	LogName     *string   `json:"log_name"`
	Description *string   `json:"description"`
	SubjectType *string   `json:"subject_type"`
	SubjectID   *uint     `json:"subject_id"`
	CauserType  *string   `json:"causer_type"`
	CauserID    *uint     `json:"causer_id"`
	Properties  JSONB     `json:"properties"`
	IPAddress   *string   `json:"ip_address"`
	UserAgent   *string   `json:"user_agent"`
	CreatedAt   time.Time `json:"created_at"`
}
