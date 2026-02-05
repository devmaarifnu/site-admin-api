package handlers

import (
	"io"
	"path/filepath"
	"strings"

	"site-admin-api/config"
	"site-admin-api/internal/middlewares"
	"site-admin-api/internal/utils"
	"site-admin-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type CDNHandler struct {
	cdnClient *utils.CDNClient
	cfg       *config.Config
}

func NewCDNHandler(cfg *config.Config) *CDNHandler {
	return &CDNHandler{
		cdnClient: utils.NewCDNClient(cfg),
		cfg:       cfg,
	}
}

// UploadRequest represents the upload request
type UploadRequest struct {
	Tag      string `form:"tag" binding:"required"`
	IsPublic bool   `form:"is_public"`
}

// UploadToCDN uploads file to CDN file server
func (h *CDNHandler) UploadToCDN(c *gin.Context) {
	// Parse multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		response.BadRequest(c, "Failed to parse form", err.Error())
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "File is required", err.Error())
		return
	}
	defer file.Close()

	// Get tag from form
	tag := c.Request.FormValue("tag")
	if tag == "" {
		response.BadRequest(c, "Tag is required", "tag field cannot be empty")
		return
	}

	// Get is_public from form (default: true for public access)
	isPublic := true
	if publicStr := c.Request.FormValue("is_public"); publicStr != "" {
		isPublic = publicStr == "true" || publicStr == "1"
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
		".svg":  true,
		".pdf":  true,
		".doc":  true,
		".docx": true,
		".xls":  true,
		".xlsx": true,
		".ppt":  true,
		".pptx": true,
	}

	if !allowedExts[ext] {
		response.BadRequest(c, "Invalid file type", "File type not allowed")
		return
	}

	// Validate file size
	maxSize := int64(h.cfg.Upload.MaxSizeImageMB) * 1024 * 1024
	if strings.Contains(tag, "document") {
		maxSize = int64(h.cfg.Upload.MaxSizeDocumentMB) * 1024 * 1024
	}

	if header.Size > maxSize {
		response.BadRequest(c, "File too large", "File size exceeds maximum allowed size")
		return
	}

	// Read file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		response.InternalServerError(c, "Failed to read file: "+err.Error())
		return
	}

	// Get uploader info
	userID := middlewares.GetUserID(c)
	userName := "Unknown"
	if user, exists := c.Get("user"); exists {
		if u, ok := user.(map[string]interface{}); ok {
			if name, ok := u["name"].(string); ok {
				userName = name
			}
		}
	}

	// Upload to CDN
	uploadResp, err := h.cdnClient.UploadFile(fileContent, header.Filename, tag, isPublic)
	if err != nil {
		response.InternalServerError(c, "Failed to upload to CDN: "+err.Error())
		return
	}

	// Prepare response
	responseData := map[string]interface{}{
		"file_id":       uploadResp.Data.FileID,
		"original_name": uploadResp.Data.OriginalName,
		"url":           uploadResp.Data.URL,
		"tag":           uploadResp.Data.Tag,
		"size":          uploadResp.Data.Size,
		"content_type":  uploadResp.Data.ContentType,
		"public":        uploadResp.Data.Public,
		"uploaded_at":   uploadResp.Data.UploadedAt,
		"uploaded_by": map[string]interface{}{
			"id":   userID,
			"name": userName,
		},
	}

	response.Created(c, "File uploaded successfully to CDN", responseData)
}

// DeleteFromCDN deletes file from CDN file server
func (h *CDNHandler) DeleteFromCDN(c *gin.Context) {
	tag := c.Param("tag")
	filename := c.Param("filename")

	if tag == "" || filename == "" {
		response.BadRequest(c, "Invalid request", "Tag and filename are required")
		return
	}

	err := h.cdnClient.DeleteFile(tag, filename)
	if err != nil {
		response.InternalServerError(c, "Failed to delete from CDN: "+err.Error())
		return
	}

	response.Success(c, "File deleted successfully from CDN", nil)
}

// GetCDNFileURL generates CDN file URL
func (h *CDNHandler) GetCDNFileURL(c *gin.Context) {
	tag := c.Query("tag")
	filename := c.Query("filename")

	if tag == "" || filename == "" {
		response.BadRequest(c, "Invalid request", "Tag and filename are required")
		return
	}

	url := h.cdnClient.GetFileURL(tag, filename)

	response.Success(c, "CDN URL generated successfully", map[string]interface{}{
		"url": url,
	})
}
