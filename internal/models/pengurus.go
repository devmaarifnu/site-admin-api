package models

import "time"

type Pengurus struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Nama          string    `gorm:"type:varchar(255);not null" json:"nama"`
	Jabatan       string    `gorm:"type:varchar(255);not null" json:"jabatan"`
	Kategori      string    `gorm:"type:enum('pimpinan_utama','bidang','sekretariat','bendahara');default:'bidang';index" json:"kategori"`
	Foto          *string   `gorm:"type:varchar(500)" json:"foto"`
	Bio           *string   `gorm:"type:text" json:"bio"`
	Email         *string   `gorm:"type:varchar(255)" json:"email"`
	Phone         *string   `gorm:"type:varchar(20)" json:"phone"`
	PeriodeMulai  int       `gorm:"type:year;not null;index" json:"periode_mulai"`
	PeriodeSelesai int      `gorm:"type:year;not null;index" json:"periode_selesai"`
	OrderNumber   int       `gorm:"default:0;index" json:"order_number"`
	IsActive      bool      `gorm:"default:true;index" json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Pengurus) TableName() string {
	return "pengurus"
}

// DTOs
type PengurusCreateRequest struct {
	Nama           string  `json:"nama" binding:"required,max=255"`
	Jabatan        string  `json:"jabatan" binding:"required,max=255"`
	Kategori       string  `json:"kategori" binding:"required,oneof=pimpinan_utama bidang sekretariat bendahara"`
	Foto           *string `json:"foto"`
	Bio            *string `json:"bio"`
	Email          *string `json:"email"`
	Phone          *string `json:"phone"`
	PeriodeMulai   int     `json:"periode_mulai" binding:"required"`
	PeriodeSelesai int     `json:"periode_selesai" binding:"required"`
	OrderNumber    int     `json:"order_number"`
	IsActive       bool    `json:"is_active"`
}

type PengurusUpdateRequest struct {
	Nama           *string `json:"nama" binding:"omitempty,max=255"`
	Jabatan        *string `json:"jabatan" binding:"omitempty,max=255"`
	Kategori       *string `json:"kategori" binding:"omitempty,oneof=pimpinan_utama bidang sekretariat bendahara"`
	Foto           *string `json:"foto"`
	Bio            *string `json:"bio"`
	Email          *string `json:"email"`
	Phone          *string `json:"phone"`
	PeriodeMulai   *int    `json:"periode_mulai"`
	PeriodeSelesai *int    `json:"periode_selesai"`
	OrderNumber    *int    `json:"order_number"`
	IsActive       *bool   `json:"is_active"`
}

type PengurusResponse struct {
	ID             uint      `json:"id"`
	Nama           string    `json:"nama"`
	Jabatan        string    `json:"jabatan"`
	Kategori       string    `json:"kategori"`
	Foto           *string   `json:"foto"`
	Bio            *string   `json:"bio"`
	Email          *string   `json:"email"`
	Phone          *string   `json:"phone"`
	PeriodeMulai   int       `json:"periode_mulai"`
	PeriodeSelesai int       `json:"periode_selesai"`
	OrderNumber    int       `json:"order_number"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
