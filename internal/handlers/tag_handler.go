package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
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
	tags, err := h.tagService.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Tags retrieved successfully", tags)
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

