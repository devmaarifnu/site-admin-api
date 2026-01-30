package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"site-admin-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	mediaService services.MediaService
}

func NewMediaHandler(mediaService services.MediaService) *MediaHandler {
	return &MediaHandler{mediaService: mediaService}
}

func (h *MediaHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)

	filters := make(map[string]interface{})
	if folder := c.Query("folder"); folder != "" {
		filters["folder"] = folder
	}
	if fileType := c.Query("file_type"); fileType != "" {
		filters["file_type"] = fileType
	}

	media, total, err := h.mediaService.GetAll(params.Page, params.Limit, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Media files retrieved successfully", media, pagination)
}

func (h *MediaHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	media, err := h.mediaService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Media file not found")
		return
	}

	response.Success(c, "Media file retrieved successfully", media)
}

func (h *MediaHandler) Upload(c *gin.Context) {
	var req models.MediaCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Get uploader ID from context
	userID, _ := c.Get("user_id")
	uploaderID := userID.(uint)

	media, err := h.mediaService.Create(&req, uploaderID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Media file uploaded successfully", media)
}

func (h *MediaHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.mediaService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Media file deleted successfully", nil)
}



