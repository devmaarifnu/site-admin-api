package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	positionService    services.OrganizationPositionService
	boardMemberService services.BoardMemberService
}

func NewOrganizationHandler(positionService services.OrganizationPositionService, boardMemberService services.BoardMemberService) *OrganizationHandler {
	return &OrganizationHandler{
		positionService:    positionService,
		boardMemberService: boardMemberService,
	}
}

// Position handlers
func (h *OrganizationHandler) GetPositions(c *gin.Context) {
	positions, err := h.positionService.GetAll(make(map[string]interface{}))
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Organization positions retrieved successfully", positions)
}

func (h *OrganizationHandler) GetPositionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	position, err := h.positionService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Position not found")
		return
	}

	response.Success(c, "Position retrieved successfully", position)
}

func (h *OrganizationHandler) CreatePosition(c *gin.Context) {
	var req models.OrganizationPositionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	position, err := h.positionService.Create(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Position created successfully", position)
}

func (h *OrganizationHandler) UpdatePosition(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.OrganizationPositionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	position, err := h.positionService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Position updated successfully", position)
}

func (h *OrganizationHandler) DeletePosition(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.positionService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Position deleted successfully", nil)
}

// Board Member handlers
func (h *OrganizationHandler) GetBoardMembers(c *gin.Context) {
	filters := make(map[string]interface{})
	if positionID := c.Query("position_id"); positionID != "" {
		filters["position_id"] = positionID
	}
	if periodStart := c.Query("period_start"); periodStart != "" {
		filters["period_start"] = periodStart
	}
	if periodEnd := c.Query("period_end"); periodEnd != "" {
		filters["period_end"] = periodEnd
	}

	members, err := h.boardMemberService.GetAll(filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Board members retrieved successfully", members)
}

func (h *OrganizationHandler) GetBoardMemberByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	member, err := h.boardMemberService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Board member not found")
		return
	}

	response.Success(c, "Board member retrieved successfully", member)
}

func (h *OrganizationHandler) CreateBoardMember(c *gin.Context) {
	var req models.BoardMemberCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	member, err := h.boardMemberService.Create(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Board member created successfully", member)
}

func (h *OrganizationHandler) UpdateBoardMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.BoardMemberUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	member, err := h.boardMemberService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Board member updated successfully", member)
}

func (h *OrganizationHandler) DeleteBoardMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.boardMemberService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Board member deleted successfully", nil)
}

// Stub methods for Pengurus
func (h *OrganizationHandler) GetPengurus(c *gin.Context) {
	response.Success(c, "Pengurus list (stub)", []interface{}{})
}

func (h *OrganizationHandler) GetPengurusByID(c *gin.Context) {
	response.Success(c, "Pengurus detail (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) CreatePengurus(c *gin.Context) {
	response.Created(c, "Pengurus created (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) UpdatePengurus(c *gin.Context) {
	response.Success(c, "Pengurus updated (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) DeletePengurus(c *gin.Context) {
	response.Success(c, "Pengurus deleted (stub)", nil)
}

func (h *OrganizationHandler) ReorderPengurus(c *gin.Context) {
	response.Success(c, "Pengurus reordered (stub)", nil)
}

// Stub methods for Departments
func (h *OrganizationHandler) GetDepartments(c *gin.Context) {
	response.Success(c, "Departments list (stub)", []interface{}{})
}

func (h *OrganizationHandler) GetDepartmentByID(c *gin.Context) {
	response.Success(c, "Department detail (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) CreateDepartment(c *gin.Context) {
	response.Created(c, "Department created (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) UpdateDepartment(c *gin.Context) {
	response.Success(c, "Department updated (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) DeleteDepartment(c *gin.Context) {
	response.Success(c, "Department deleted (stub)", nil)
}

// Stub methods for Editorial
func (h *OrganizationHandler) GetEditorialTeam(c *gin.Context) {
	response.Success(c, "Editorial team (stub)", []interface{}{})
}

func (h *OrganizationHandler) UpdateEditorialTeam(c *gin.Context) {
	response.Success(c, "Editorial team updated (stub)", nil)
}

func (h *OrganizationHandler) GetEditorialCouncil(c *gin.Context) {
	response.Success(c, "Editorial council (stub)", []interface{}{})
}

func (h *OrganizationHandler) UpdateEditorialCouncil(c *gin.Context) {
	response.Success(c, "Editorial council updated (stub)", nil)
}
