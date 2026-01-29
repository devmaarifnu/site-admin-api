package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"site-admin-api/pkg/response"
)

// PaginationParams holds pagination parameters
type PaginationParams struct {
	Page     int
	Limit    int
	Offset   int
	SortBy   string
	SortDir  string
}

// GetPaginationParams extracts pagination parameters from request
func GetPaginationParams(c *gin.Context, defaultLimit, maxLimit int) PaginationParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(defaultLimit)))
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortDir := c.DefaultQuery("sort_dir", "desc")

	// Validate page
	if page < 1 {
		page = 1
	}

	// Validate limit
	if limit < 1 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	// Validate sort direction
	if sortDir != "asc" && sortDir != "desc" {
		sortDir = "desc"
	}

	offset := (page - 1) * limit

	return PaginationParams{
		Page:    page,
		Limit:   limit,
		Offset:  offset,
		SortBy:  sortBy,
		SortDir: sortDir,
	}
}

// CalculatePaginationMeta calculates pagination metadata
func CalculatePaginationMeta(page, limit int, totalItems int64) response.PaginationMeta {
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	if totalPages < 1 {
		totalPages = 1
	}

	return response.PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   totalItems,
		ItemsPerPage: limit,
		HasNext:      page < totalPages,
		HasPrev:      page > 1,
	}
}

// GetSearchQuery extracts search query from request
func GetSearchQuery(c *gin.Context) string {
	return c.Query("search")
}

// GetFilterParams extracts filter parameters from request
func GetFilterParams(c *gin.Context) map[string]interface{} {
	filters := make(map[string]interface{})

	// Common filters
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if categoryID := c.Query("category_id"); categoryID != "" {
		filters["category_id"] = categoryID
	}
	if isFeatured := c.Query("is_featured"); isFeatured != "" {
		filters["is_featured"] = isFeatured == "true" || isFeatured == "1"
	}
	if isActive := c.Query("is_active"); isActive != "" {
		filters["is_active"] = isActive == "true" || isActive == "1"
	}
	if isPublic := c.Query("is_public"); isPublic != "" {
		filters["is_public"] = isPublic == "true" || isPublic == "1"
	}

	return filters
}
