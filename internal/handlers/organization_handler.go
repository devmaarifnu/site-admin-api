package handlers

import (
	"site-admin-api/internal/models"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	positionService         services.OrganizationPositionService
	boardMemberService      services.BoardMemberService
	pengurusService         services.PengurusService
	departmentService       services.DepartmentService
	editorialTeamService    services.EditorialTeamService
	editorialCouncilService services.EditorialCouncilService
}

func NewOrganizationHandler(
	positionService services.OrganizationPositionService,
	boardMemberService services.BoardMemberService,
	pengurusService services.PengurusService,
	departmentService services.DepartmentService,
	editorialTeamService services.EditorialTeamService,
	editorialCouncilService services.EditorialCouncilService,
) *OrganizationHandler {
	return &OrganizationHandler{
		positionService:         positionService,
		boardMemberService:      boardMemberService,
		pengurusService:         pengurusService,
		departmentService:       departmentService,
		editorialTeamService:    editorialTeamService,
		editorialCouncilService: editorialCouncilService,
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
	filters := make(map[string]interface{})
	if kategori := c.Query("kategori"); kategori != "" {
		filters["kategori"] = kategori
	}
	if isActive := c.Query("is_active"); isActive != "" {
		filters["is_active"] = isActive
	}

	pengurusList, err := h.pengurusService.GetAll(filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Pengurus retrieved successfully", pengurusList)
}

func (h *OrganizationHandler) GetPengurusByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	pengurus, err := h.pengurusService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Pengurus not found")
		return
	}

	response.Success(c, "Pengurus retrieved successfully", pengurus)
}

func (h *OrganizationHandler) CreatePengurus(c *gin.Context) {
	var req models.PengurusCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	pengurus, err := h.pengurusService.Create(&req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "Pengurus created successfully", pengurus)
}

func (h *OrganizationHandler) UpdatePengurus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.PengurusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	pengurus, err := h.pengurusService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Pengurus updated successfully", pengurus)
}

func (h *OrganizationHandler) DeletePengurus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	if err := h.pengurusService.Delete(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Pengurus deleted successfully", nil)
}

func (h *OrganizationHandler) ReorderPengurus(c *gin.Context) {
	var req []map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	if err := h.pengurusService.Reorder(req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Pengurus reordered successfully", nil)
}

// Stub methods for Departments
func (h *OrganizationHandler) GetDepartments(c *gin.Context) {
	filters := make(map[string]interface{})
	if isActive := c.Query("is_active"); isActive != "" {
		filters["is_active"] = isActive
	}

	departments, err := h.departmentService.GetAll(filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Departments retrieved successfully", departments)
}

func (h *OrganizationHandler) GetDepartmentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	dept, err := h.departmentService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Department not found")
		return
	}

	response.Success(c, "Department retrieved successfully", dept)
}

func (h *OrganizationHandler) CreateDepartment(c *gin.Context) {
	response.Created(c, "Department created (stub)", map[string]interface{}{})
}

func (h *OrganizationHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.DepartmentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	dept, err := h.departmentService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Department updated successfully", dept)
}

func (h *OrganizationHandler) DeleteDepartment(c *gin.Context) {
	response.Success(c, "Department deleted (stub)", nil)
}

// Stub methods for Editorial
func (h *OrganizationHandler) GetEditorialTeam(c *gin.Context) {
	filters := make(map[string]interface{})
	if roleType := c.Query("role_type"); roleType != "" {
		filters["role_type"] = roleType
	}
	if isActive := c.Query("is_active"); isActive != "" {
		filters["is_active"] = isActive
	}

	members, err := h.editorialTeamService.GetAll(filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Editorial team retrieved successfully", members)
}

func (h *OrganizationHandler) UpdateEditorialTeam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.EditorialTeamUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	member, err := h.editorialTeamService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Editorial team member updated successfully", member)
}

func (h *OrganizationHandler) GetEditorialCouncil(c *gin.Context) {
	filters := make(map[string]interface{})
	if isActive := c.Query("is_active"); isActive != "" {
		filters["is_active"] = isActive
	}

	members, err := h.editorialCouncilService.GetAll(filters)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Editorial council retrieved successfully", members)
}

func (h *OrganizationHandler) UpdateEditorialCouncil(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", err.Error())
		return
	}

	var req models.EditorialCouncilUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	member, err := h.editorialCouncilService.Update(uint(id), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Editorial council member updated successfully", member)
}
