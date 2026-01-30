package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	filters := make(map[string]interface{})
	if categoryType := c.Query("type"); categoryType != "" {
		filters["type"] = categoryType
	}

	categories, err := h.categoryService.GetAll(filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Categories retrieved successfully", categories)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	category, err := h.categoryService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Category not found")
		return
	}

	response.Success(c, "Category retrieved successfully", category)
}

func (h *CategoryHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	category, err := h.categoryService.GetBySlug(slug)
	if err != nil {
		response.NotFound(c, "Category not found")
		return
	}

	response.Success(c, "Category retrieved successfully", category)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req models.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	category, err := h.categoryService.Create(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Category created successfully", category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	category, err := h.categoryService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Category updated successfully", category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.categoryService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Category deleted successfully", nil)
}
