package handlers

import (
	"site-admin-api/internal/middlewares"
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/internal/utils"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	documentService services.DocumentService
}

func NewDocumentHandler(documentService services.DocumentService) *DocumentHandler {
	return &DocumentHandler{documentService: documentService}
}

func (h *DocumentHandler) GetAll(c *gin.Context) {
	params := utils.GetPaginationParams(c, 20, 100)
	search := c.Query("search")

	filters := make(map[string]interface{})
	if categoryID := c.Query("category_id"); categoryID != "" {
		filters["category_id"] = categoryID
	}
	if docType := c.Query("type"); docType != "" {
		filters["type"] = docType
	}
	if year := c.Query("year"); year != "" {
		filters["year"] = year
	}

	documents, total, err := h.documentService.GetAll(params.Page, params.Limit, search, filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	pagination := utils.CalculatePaginationMeta(params.Page, params.Limit, total)
	response.SuccessWithPagination(c, "Documents retrieved successfully", documents, pagination)
}

func (h *DocumentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	document, err := h.documentService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Document not found")
		return
	}

	response.Success(c, "Document retrieved successfully", document)
}

func (h *DocumentHandler) Create(c *gin.Context) {
	var req models.DocumentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Get uploader ID from context (set by auth middleware)
	uploaderID := middlewares.GetUserID(c)

	document, err := h.documentService.Create(&req, uploaderID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Document created successfully", document)
}

func (h *DocumentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.DocumentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	document, err := h.documentService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Document updated successfully", document)
}

func (h *DocumentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.documentService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Document deleted successfully", nil)
}

func (h *DocumentHandler) IncrementDownloads(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.documentService.IncrementDownloads(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Downloads incremented successfully", nil)
}

func (h *DocumentHandler) GetStats(c *gin.Context) {
	stats, err := h.documentService.GetStats()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Document statistics retrieved successfully", stats)
}

func (h *DocumentHandler) ReplaceFile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		CategoryID  *uint   `json:"category_id"`
		FileName    *string `json:"file_name"`
		FileURL     string  `json:"file_url" binding:"required"`
		FileType    *string `json:"file_type"`
		FileSize    *uint64 `json:"file_size"`
		MimeType    *string `json:"mime_type"`
		IsPublic    *bool   `json:"is_public"`
		Status      *string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	updateReq := &models.DocumentUpdateRequest{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		FileName:    req.FileName,
		FileURL:     &req.FileURL,
		FileType:    req.FileType,
		FileSize:    req.FileSize,
		MimeType:    req.MimeType,
		IsPublic:    req.IsPublic,
		Status:      req.Status,
	}

	document, err := h.documentService.Update(uint(id), updateReq)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Document file replaced successfully", document)
}

func (h *DocumentHandler) Upload(c *gin.Context) {
	var req models.DocumentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Get uploader ID from context (set by auth middleware)
	uploaderID := middlewares.GetUserID(c)

	document, err := h.documentService.Create(&req, uploaderID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Document uploaded successfully", document)
}
