package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HeroSlideHandler struct {
	heroSlideService services.HeroSlideService
}

func NewHeroSlideHandler(heroSlideService services.HeroSlideService) *HeroSlideHandler {
	return &HeroSlideHandler{heroSlideService: heroSlideService}
}

func (h *HeroSlideHandler) GetAll(c *gin.Context) {
	slides, err := h.heroSlideService.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Hero slides retrieved successfully", slides)
}

func (h *HeroSlideHandler) GetActive(c *gin.Context) {
	slides, err := h.heroSlideService.GetActive()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Active hero slides retrieved successfully", slides)
}

func (h *HeroSlideHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	slide, err := h.heroSlideService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Hero slide not found")
		return
	}

	response.Success(c, "Hero slide retrieved successfully", slide)
}

func (h *HeroSlideHandler) Create(c *gin.Context) {
	var req models.HeroSlideCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	slide, err := h.heroSlideService.Create(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Hero slide created successfully", slide)
}

func (h *HeroSlideHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.HeroSlideUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	slide, err := h.heroSlideService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Hero slide updated successfully", slide)
}

func (h *HeroSlideHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.heroSlideService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Hero slide deleted successfully", nil)
}

func (h *HeroSlideHandler) Reorder(c *gin.Context) {
	var req models.HeroSlideReorderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.heroSlideService.Reorder(&req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Hero slides reordered successfully", nil)
}

func (h *HeroSlideHandler) ToggleStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	slide, err := h.heroSlideService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Hero slide not found")
		return
	}

	// Toggle status
	newStatus := !slide.IsActive
	updateReq := &models.HeroSlideUpdateRequest{
		IsActive: &newStatus,
	}

	updatedSlide, err := h.heroSlideService.Update(uint(id), updateReq)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Hero slide status toggled successfully", updatedSlide)
}
