package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"site-admin-api/internal/utils"
)

type ContactMessageHandler struct {
	contactService services.ContactMessageService
}

func NewContactMessageHandler(contactService services.ContactMessageService) *ContactMessageHandler {
	return &ContactMessageHandler{contactService: contactService}
}

func (h *ContactMessageHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if priority := c.Query("priority"); priority != "" {
		filters["priority"] = priority
	}

	messages, total, err := h.contactService.GetAll(params.Page, params.Limit, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Contact messages retrieved successfully", messages, pagination)
}

func (h *ContactMessageHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	message, err := h.contactService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Contact message not found")
		return
	}

	response.Success(c, "Contact message retrieved successfully", message)
}

func (h *ContactMessageHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req struct {
		Status string `json:"status" binding:"required,oneof=new read resolved archived"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	message, err := h.contactService.UpdateStatus(uint(id), req.Status)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Contact message status updated successfully", message)
}

func (h *ContactMessageHandler) Reply(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req struct {
		ReplyMessage string `json:"reply_message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Get replier ID from context
	userID, _ := c.Get("user_id")
	replierID := userID.(uint)

	message, err := h.contactService.Reply(uint(id), req.ReplyMessage, replierID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Reply sent successfully", message)
}

func (h *ContactMessageHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.contactService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Contact message deleted successfully", nil)
}
