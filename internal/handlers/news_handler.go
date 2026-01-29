package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/internal/utils"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	newsService services.NewsService
}

func NewNewsHandler(newsService services.NewsService) *NewsHandler {
	return &NewsHandler{newsService: newsService}
}

func (h *NewsHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)
	search := c.Query("search")

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if categoryID := c.Query("category_id"); categoryID != "" {
		filters["category_id"] = categoryID
	}
	if isFeatured := c.Query("is_featured"); isFeatured != "" {
		filters["is_featured"] = isFeatured
	}

	news, total, err := h.newsService.GetAll(params.Page, params.Limit, search, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "News articles retrieved successfully", news, pagination)
}

func (h *NewsHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	news, err := h.newsService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, "News article retrieved successfully", news)
}

func (h *NewsHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	news, err := h.newsService.GetBySlug(slug)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, "News article retrieved successfully", news)
}

func (h *NewsHandler) Create(c *gin.Context) {
	var req models.NewsArticleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	userID := c.GetUint("user_id")
	news, err := h.newsService.Create(&req, userID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "News article created successfully", news)
}

func (h *NewsHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.NewsArticleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	news, err := h.newsService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "News article updated successfully", news)
}

func (h *NewsHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.newsService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "News article deleted successfully", nil)
}

func (h *NewsHandler) IncrementViews(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.newsService.IncrementViews(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Views incremented successfully", nil)
}

func (h *NewsHandler) GetFeatured(c *gin.Context) {
	limit := 5
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil {
			limit = parsed
		}
	}

	news, err := h.newsService.GetFeatured(limit)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Featured news retrieved successfully", news)
}

// Additional methods for routes compatibility
func (h *NewsHandler) Publish(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	req := &models.NewsArticleUpdateRequest{
		Status: stringPtr("published"),
	}

	news, err := h.newsService.Update(uint(id), req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "News article published successfully", news)
}

func (h *NewsHandler) Archive(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	req := &models.NewsArticleUpdateRequest{
		Status: stringPtr("archived"),
	}

	news, err := h.newsService.Update(uint(id), req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "News article archived successfully", news)
}

func (h *NewsHandler) ToggleFeatured(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	news, err := h.newsService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	req := &models.NewsArticleUpdateRequest{
		IsFeatured: boolPtr(!news.IsFeatured),
	}

	updatedNews, err := h.newsService.Update(uint(id), req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Featured status toggled successfully", updatedNews)
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
