package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type SettingRepository interface {
	FindAll(group string) ([]models.Setting, error)
	FindByKey(key string) (*models.Setting, error)
	FindPublic() ([]models.Setting, error)
	FindByGroup(group string) ([]models.Setting, error)
	Create(setting *models.Setting) error
	Update(setting *models.Setting) error
	Delete(key string) error
	Upsert(setting *models.Setting) error
}

type settingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

func (r *settingRepository) FindAll(group string) ([]models.Setting, error) {
	var settings []models.Setting
	query := r.db.Model(&models.Setting{})

	if group != "" {
		query = query.Where("setting_group = ?", group)
	}

	if err := query.Order("setting_group ASC, setting_key ASC").Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

func (r *settingRepository) FindByKey(key string) (*models.Setting, error) {
	var setting models.Setting
	if err := r.db.Where("setting_key = ?", key).First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r *settingRepository) FindPublic() ([]models.Setting, error) {
	var settings []models.Setting
	if err := r.db.Where("is_public = ?", true).Order("setting_group ASC, setting_key ASC").Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

func (r *settingRepository) FindByGroup(group string) ([]models.Setting, error) {
	var settings []models.Setting
	if err := r.db.Where("setting_group = ?", group).Order("setting_key ASC").Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

func (r *settingRepository) Create(setting *models.Setting) error {
	return r.db.Create(setting).Error
}

func (r *settingRepository) Update(setting *models.Setting) error {
	return r.db.Save(setting).Error
}

func (r *settingRepository) Delete(key string) error {
	return r.db.Where("setting_key = ?", key).Delete(&models.Setting{}).Error
}

func (r *settingRepository) Upsert(setting *models.Setting) error {
	var existing models.Setting
	err := r.db.Where("setting_key = ?", setting.SettingKey).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(setting).Error
	} else if err != nil {
		return err
	}

	setting.ID = existing.ID
	return r.db.Save(setting).Error
}
