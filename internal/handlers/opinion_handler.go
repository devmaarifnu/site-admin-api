package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"site-admin-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OpinionHandler struct {
	opinionService services.OpinionService
}

func NewOpinionHandler(opinionService services.OpinionService) *OpinionHandler {
	return &OpinionHandler{opinionService: opinionService}
}

func (h *OpinionHandler) GetAll(c *gin.Context) {
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

	opinions, total, err := h.opinionService.GetAll(params.Page, params.Limit, search, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Opinion articles retrieved successfully", opinions, pagination)
}

func (h *OpinionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	opinion, err := h.opinionService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Opinion article not found")
		return
	}

	response.Success(c, "Opinion article retrieved successfully", opinion)
}

func (h *OpinionHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	opinion, err := h.opinionService.GetBySlug(slug)
	if err != nil {
		response.NotFound(c, "Opinion article not found")
		return
	}

	response.Success(c, "Opinion article retrieved successfully", opinion)
}

func (h *OpinionHandler) Create(c *gin.Context) {
	var req models.OpinionArticleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Get author ID from context (set by auth middleware)
	userID, _ := c.Get("user_id")
	authorID := userID.(uint)

	opinion, err := h.opinionService.Create(&req, authorID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Opinion article created successfully", opinion)
}

func (h *OpinionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.OpinionArticleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	opinion, err := h.opinionService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Opinion article updated successfully", opinion)
}

func (h *OpinionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.opinionService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Opinion article deleted successfully", nil)
}

func (h *OpinionHandler) Publish(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	req := &models.OpinionArticleUpdateRequest{
		Status: stringPtr("published"),
	}

	opinion, err := h.opinionService.Update(uint(id), req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Opinion article published successfully", opinion)
}

func (h *OpinionHandler) IncrementViews(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.opinionService.IncrementViews(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Views incremented successfully", nil)
}

func (h *OpinionHandler) GetFeatured(c *gin.Context) {
	limit := 5
	if l := c.Query("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil {
			limit = parsedLimit
		}
	}

	opinions, err := h.opinionService.GetFeatured(limit)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Featured opinion articles retrieved successfully", opinions)
}
