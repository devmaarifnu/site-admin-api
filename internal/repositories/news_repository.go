package repositories

import (
	"site-admin-api/internal/models"

	"gorm.io/gorm"
)

type NewsRepository interface {
	FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.NewsArticle, int64, error)
	FindByID(id uint) (*models.NewsArticle, error)
	FindBySlug(slug string) (*models.NewsArticle, error)
	Create(news *models.NewsArticle) error
	Update(news *models.NewsArticle) error
	Delete(id uint) error
	IncrementViews(id uint) error
	FindFeatured(limit int) ([]models.NewsArticle, error)
	FindByCategory(categoryID uint, page, limit int) ([]models.NewsArticle, int64, error)
	FindByTag(tagID uint, page, limit int) ([]models.NewsArticle, int64, error)
	UpdateTags(newsID uint, tagIDs []uint) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) FindAll(page, limit int, search string, filters map[string]interface{}) ([]models.NewsArticle, int64, error) {
	var news []models.NewsArticle
	var total int64

	query := r.db.Model(&models.NewsArticle{}).Preload("Category").Preload("Author").Preload("Tags")

	// Search
	if search != "" {
		query = query.Where("title LIKE ? OR excerpt LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Filters
	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&news).Error; err != nil {
		return nil, 0, err
	}

	return news, total, nil
}

func (r *newsRepository) FindByID(id uint) (*models.NewsArticle, error) {
	var news models.NewsArticle
	if err := r.db.Preload("Category").Preload("Author").Preload("Tags").First(&news, id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) FindBySlug(slug string) (*models.NewsArticle, error) {
	var news models.NewsArticle
	if err := r.db.Preload("Category").Preload("Author").Preload("Tags").Where("slug = ?", slug).First(&news).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) Create(news *models.NewsArticle) error {
	return r.db.Create(news).Error
}

func (r *newsRepository) Update(news *models.NewsArticle) error {
	return r.db.Save(news).Error
}

func (r *newsRepository) Delete(id uint) error {
	return r.db.Delete(&models.NewsArticle{}, id).Error
}

func (r *newsRepository) IncrementViews(id uint) error {
	return r.db.Model(&models.NewsArticle{}).Where("id = ?", id).UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

func (r *newsRepository) FindFeatured(limit int) ([]models.NewsArticle, error) {
	var news []models.NewsArticle
	if err := r.db.Where("is_featured = ? AND status = ?", true, "published").
		Preload("Category").Preload("Author").Preload("Tags").
		Order("published_at DESC").Limit(limit).Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) FindByCategory(categoryID uint, page, limit int) ([]models.NewsArticle, int64, error) {
	var news []models.NewsArticle
	var total int64

	query := r.db.Model(&models.NewsArticle{}).Where("category_id = ?", categoryID).
		Preload("Category").Preload("Author").Preload("Tags")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&news).Error; err != nil {
		return nil, 0, err
	}

	return news, total, nil
}

func (r *newsRepository) FindByTag(tagID uint, page, limit int) ([]models.NewsArticle, int64, error) {
	var news []models.NewsArticle
	var total int64

	query := r.db.Model(&models.NewsArticle{}).
		Joins("JOIN news_tags ON news_tags.news_article_id = news_articles.id").
		Where("news_tags.tag_id = ?", tagID).
		Preload("Category").Preload("Author").Preload("Tags")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("news_articles.created_at DESC").Offset(offset).Limit(limit).Find(&news).Error; err != nil {
		return nil, 0, err
	}

	return news, total, nil
}

func (r *newsRepository) UpdateTags(newsID uint, tagIDs []uint) error {
	var news models.NewsArticle
	if err := r.db.First(&news, newsID).Error; err != nil {
		return err
	}

	// Clear existing tags
	if err := r.db.Model(&news).Association("Tags").Clear(); err != nil {
		return err
	}

	// Add new tags
	if len(tagIDs) > 0 {
		var tags []models.Tag
		if err := r.db.Find(&tags, tagIDs).Error; err != nil {
			return err
		}
		if err := r.db.Model(&news).Association("Tags").Append(tags); err != nil {
			return err
		}
	}

	return nil
}
