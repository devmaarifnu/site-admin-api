package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"site-admin-api/internal/services"
	"site-admin-api/internal/utils"
	"site-admin-api/pkg/response"
)

type ActivityLogHandler struct {
	activityLogService services.ActivityLogService
}

func NewActivityLogHandler(activityLogService services.ActivityLogService) *ActivityLogHandler {
	return &ActivityLogHandler{activityLogService: activityLogService}
}

func (h *ActivityLogHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)

	filters := make(map[string]interface{})
	if userID := c.Query("user_id"); userID != "" {
		filters["user_id"] = userID
	}
	if action := c.Query("action"); action != "" {
		filters["action"] = action
	}
	if entityType := c.Query("entity_type"); entityType != "" {
		filters["entity_type"] = entityType
	}

	logs, total, err := h.activityLogService.GetAll(params.Page, params.Limit, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Activity logs retrieved successfully", logs, pagination)
}

func (h *ActivityLogHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	log, err := h.activityLogService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Activity log not found")
		return
	}

	response.Success(c, "Activity log retrieved successfully", log)
}
