package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"site-admin-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventFlyerHandler struct {
	eventFlyerService services.EventFlyerService
}

func NewEventFlyerHandler(eventFlyerService services.EventFlyerService) *EventFlyerHandler {
	return &EventFlyerHandler{eventFlyerService: eventFlyerService}
}

func (h *EventFlyerHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)
	search := c.Query("search")

	flyers, total, err := h.eventFlyerService.GetAll(params.Page, params.Limit, search)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Event flyers retrieved successfully", flyers, pagination)
}

func (h *EventFlyerHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	flyer, err := h.eventFlyerService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Event flyer not found")
		return
	}

	response.Success(c, "Event flyer retrieved successfully", flyer)
}

func (h *EventFlyerHandler) Create(c *gin.Context) {
	var req models.EventFlyerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Get uploader ID from context
	userID, _ := c.Get("user_id")
	uploaderID := userID.(uint)

	flyer, err := h.eventFlyerService.Create(&req, uploaderID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Event flyer created successfully", flyer)
}

func (h *EventFlyerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.eventFlyerService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Event flyer deleted successfully", nil)
}
