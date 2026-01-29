package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"site-admin-api/internal/utils"
)

type NotificationHandler struct {
	notificationService services.NotificationService
}

func NewNotificationHandler(notificationService services.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

func (h *NotificationHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)

	// Get user ID from context
	userID, _ := c.Get("user_id")

	notifications, total, err := h.notificationService.GetAll(userID.(uint), params.Page, params.Limit)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Notifications retrieved successfully", notifications, pagination)
}

func (h *NotificationHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	notification, err := h.notificationService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Notification not found")
		return
	}

	response.Success(c, "Notification retrieved successfully", notification)
}

func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	notification, err := h.notificationService.MarkAsRead(uint(id))
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Notification marked as read", notification)
}

func (h *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	// Get user ID from context
	userID, _ := c.Get("user_id")

	if err := h.notificationService.MarkAllAsRead(userID.(uint)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "All notifications marked as read", nil)
}

func (h *NotificationHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.notificationService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Notification deleted successfully", nil)
}
