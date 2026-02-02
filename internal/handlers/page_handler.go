package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	pageService services.PageService
}

func NewPageHandler(pageService services.PageService) *PageHandler {
	return &PageHandler{pageService: pageService}
}

func (h *PageHandler) GetAll(c *gin.Context) {
	pages, err := h.pageService.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Pages retrieved successfully", pages)
}

func (h *PageHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	page, err := h.pageService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Page not found")
		return
	}

	response.Success(c, "Page retrieved successfully", page)
}

func (h *PageHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	page, err := h.pageService.GetBySlug(slug)
	if err != nil {
		response.NotFound(c, "Page not found")
		return
	}

	response.Success(c, "Page retrieved successfully", page)
}

func (h *PageHandler) Update(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		response.BadRequest(c, "Slug is required", "Slug parameter is empty")
		return
	}

	var req models.PageUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	page, err := h.pageService.UpdateBySlug(slug, &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Page updated successfully", page)
}
