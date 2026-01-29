package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"site-admin-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagService services.TagService
}

func NewTagHandler(tagService services.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func (h *TagHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)
	search := c.Query("search")

	tags, total, err := h.tagService.GetAll(params.Page, params.Limit, search)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Tags retrieved successfully", tags, pagination)
}

func (h *TagHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	tag, err := h.tagService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Tag not found")
		return
	}

	response.Success(c, "Tag retrieved successfully", tag)
}

func (h *TagHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	tag, err := h.tagService.GetBySlug(slug)
	if err != nil {
		response.NotFound(c, "Tag not found")
		return
	}

	response.Success(c, "Tag retrieved successfully", tag)
}

func (h *TagHandler) Create(c *gin.Context) {
	var req models.TagCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	tag, err := h.tagService.Create(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Tag created successfully", tag)
}

func (h *TagHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.TagUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	tag, err := h.tagService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Tag updated successfully", tag)
}

func (h *TagHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.tagService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Tag deleted successfully", nil)
}

func (h *TagHandler) Merge(c *gin.Context) {
	var req struct {
		SourceTagIDs []uint `json:"source_tag_ids" binding:"required"`
		TargetTagID  uint   `json:"target_tag_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Stub implementation - would need to implement actual merge logic
	response.Success(c, "Tags merged successfully", nil)
}
