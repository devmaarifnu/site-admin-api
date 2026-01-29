package models

import "time"

type HeroSlide struct {
	ID                   uint       `gorm:"primaryKey" json:"id"`
	Title                string     `gorm:"type:varchar(500);not null" json:"title"`
	Description          *string    `gorm:"type:text" json:"description"`
	Image                string     `gorm:"type:varchar(500);not null" json:"image"`
	CTALabel             *string    `gorm:"type:varchar(100)" json:"cta_label"`
	CTAHref              *string    `gorm:"type:varchar(500)" json:"cta_href"`
	CTASecondaryLabel    *string    `gorm:"type:varchar(100)" json:"cta_secondary_label"`
	CTASecondaryHref     *string    `gorm:"type:varchar(500)" json:"cta_secondary_href"`
	OrderNumber          int        `gorm:"default:0;index" json:"order_number"`
	IsActive             bool       `gorm:"default:true;index" json:"is_active"`
	StartDate            *time.Time `json:"start_date"`
	EndDate              *time.Time `json:"end_date"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

func (HeroSlide) TableName() string {
	return "hero_slides"
}

// DTOs
type HeroSlideCreateRequest struct {
	Title             string     `json:"title" binding:"required,max=500"`
	Description       *string    `json:"description"`
	Image             string     `json:"image" binding:"required"`
	CTALabel          *string    `json:"cta_label"`
	CTAHref           *string    `json:"cta_href"`
	CTASecondaryLabel *string    `json:"cta_secondary_label"`
	CTASecondaryHref  *string    `json:"cta_secondary_href"`
	OrderNumber       int        `json:"order_number"`
	IsActive          bool       `json:"is_active"`
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
}

type HeroSlideUpdateRequest struct {
	Title             *string    `json:"title" binding:"omitempty,max=500"`
	Description       *string    `json:"description"`
	Image             *string    `json:"image"`
	CTALabel          *string    `json:"cta_label"`
	CTAHref           *string    `json:"cta_href"`
	CTASecondaryLabel *string    `json:"cta_secondary_label"`
	CTASecondaryHref  *string    `json:"cta_secondary_href"`
	OrderNumber       *int       `json:"order_number"`
	IsActive          *bool      `json:"is_active"`
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
}

type HeroSlideResponse struct {
	ID                uint       `json:"id"`
	Title             string     `json:"title"`
	Subtitle          *string    `json:"subtitle"` // alias for Description
	Description       *string    `json:"description"`
	Image             string     `json:"image"`
	ImageURL          string     `json:"image_url"` // alias for Image
	CTALabel          *string    `json:"cta_label"`
	CTAHref           *string    `json:"cta_href"`
	LinkURL           *string    `json:"link_url"` // alias for CTAHref
	LinkText          *string    `json:"link_text"` // alias for CTALabel
	CTASecondaryLabel *string    `json:"cta_secondary_label"`
	CTASecondaryHref  *string    `json:"cta_secondary_href"`
	OrderNumber       int        `json:"order_number"`
	DisplayOrder      int        `json:"display_order"` // alias for OrderNumber
	IsActive          bool       `json:"is_active"`
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type HeroSlideReorderRequest struct {
	SlideOrders []struct {
		ID    uint `json:"id" binding:"required"`
		Order int  `json:"order" binding:"required"`
	} `json:"slide_orders" binding:"required"`
}
