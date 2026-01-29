package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type OpinionRepository interface {
	FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.OpinionArticle, int64, error)
	FindByID(id uint) (*models.OpinionArticle, error)
	FindBySlug(slug string) (*models.OpinionArticle, error)
	Create(opinion *models.OpinionArticle) error
	Update(opinion *models.OpinionArticle) error
	Delete(id uint) error
	IncrementViews(id uint) error
	FindFeatured(limit int) ([]models.OpinionArticle, error)
	FindByTag(tagID uint, page, limit int) ([]models.OpinionArticle, int64, error)
	UpdateTags(opinionID uint, tagIDs []uint) error
}

type opinionRepository struct {
	db *gorm.DB
}

func NewOpinionRepository(db *gorm.DB) OpinionRepository {
	return &opinionRepository{db: db}
}

func (r *opinionRepository) FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.OpinionArticle, int64, error) {
	var opinions []models.OpinionArticle
	var total int64

	query := r.db.Model(&models.OpinionArticle{}).Preload("Creator").Preload("Tags")

	if search != "" {
		query = query.Where("title LIKE ? OR excerpt LIKE ? OR content LIKE ? OR author_name LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&opinions).Error; err != nil {
		return nil, 0, err
	}

	return opinions, total, nil
}

func (r *opinionRepository) FindByID(id uint) (*models.OpinionArticle, error) {
	var opinion models.OpinionArticle
	if err := r.db.Preload("Creator").Preload("Tags").First(&opinion, id).Error; err != nil {
		return nil, err
	}
	return &opinion, nil
}

func (r *opinionRepository) FindBySlug(slug string) (*models.OpinionArticle, error) {
	var opinion models.OpinionArticle
	if err := r.db.Preload("Creator").Preload("Tags").Where("slug = ?", slug).First(&opinion).Error; err != nil {
		return nil, err
	}
	return &opinion, nil
}

func (r *opinionRepository) Create(opinion *models.OpinionArticle) error {
	return r.db.Create(opinion).Error
}

func (r *opinionRepository) Update(opinion *models.OpinionArticle) error {
	return r.db.Save(opinion).Error
}

func (r *opinionRepository) Delete(id uint) error {
	return r.db.Delete(&models.OpinionArticle{}, id).Error
}

func (r *opinionRepository) IncrementViews(id uint) error {
	return r.db.Model(&models.OpinionArticle{}).Where("id = ?", id).UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

func (r *opinionRepository) FindFeatured(limit int) ([]models.OpinionArticle, error) {
	var opinions []models.OpinionArticle
	if err := r.db.Where("is_featured = ? AND status = ?", true, "published").
		Preload("Creator").Preload("Tags").
		Order("published_at DESC").Limit(limit).Find(&opinions).Error; err != nil {
		return nil, err
	}
	return opinions, nil
}

func (r *opinionRepository) FindByTag(tagID uint, page, limit int) ([]models.OpinionArticle, int64, error) {
	var opinions []models.OpinionArticle
	var total int64

	query := r.db.Model(&models.OpinionArticle{}).
		Joins("JOIN opinion_tags ON opinion_tags.opinion_article_id = opinion_articles.id").
		Where("opinion_tags.tag_id = ?", tagID).
		Preload("Creator").Preload("Tags")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("opinion_articles.created_at DESC").Offset(offset).Limit(limit).Find(&opinions).Error; err != nil {
		return nil, 0, err
	}

	return opinions, total, nil
}

func (r *opinionRepository) UpdateTags(opinionID uint, tagIDs []uint) error {
	var opinion models.OpinionArticle
	if err := r.db.First(&opinion, opinionID).Error; err != nil {
		return err
	}

	if err := r.db.Model(&opinion).Association("Tags").Clear(); err != nil {
		return err
	}

	if len(tagIDs) > 0 {
		var tags []models.Tag
		if err := r.db.Find(&tags, tagIDs).Error; err != nil {
			return err
		}
		if err := r.db.Model(&opinion).Association("Tags").Append(tags); err != nil {
			return err
		}
	}

	return nil
}
