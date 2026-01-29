package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type SettingHandler struct {
	settingService services.SettingService
}

func NewSettingHandler(settingService services.SettingService) *SettingHandler {
	return &SettingHandler{settingService: settingService}
}

func (h *SettingHandler) GetAll(c *gin.Context) {
	group := c.Query("group")

	settings, err := h.settingService.GetAll(group)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Settings retrieved successfully", settings)
}

func (h *SettingHandler) GetByKey(c *gin.Context) {
	key := c.Param("key")

	setting, err := h.settingService.GetByKey(key)
	if err != nil {
		response.NotFound(c, "Setting not found")
		return
	}

	response.Success(c, "Setting retrieved successfully", setting)
}

func (h *SettingHandler) Update(c *gin.Context) {
	key := c.Param("key")

	var req struct {
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	setting, err := h.settingService.Update(key, req.Value)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Setting updated successfully", setting)
}

func (h *SettingHandler) BulkUpdate(c *gin.Context) {
	var req struct {
		Settings []models.SettingUpdateRequest `json:"settings" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.settingService.BulkUpdate(req.Settings); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Settings updated successfully", nil)
}
