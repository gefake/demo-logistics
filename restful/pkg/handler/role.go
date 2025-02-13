package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// CreateRole godoc
//
// @Summary Create a new role
// @Description Create a new role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body database.Role true "Role"
// @Success 201 {object} errorResponse "Role created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/role [post]
func (h *Handler) CreateRole(c *gin.Context) {
	var role database.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.RoleRepository.CreateRole(role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetRoleByID godoc
//
// @Summary Get role by ID
// @Description Get role by ID
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} database.Role
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/role/{id} [get]
func (h *Handler) GetRoleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid role ID")
		return
	}

	role, err := h.Services.RoleRepository.GetRoleByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Role not found")
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole godoc
//
// @Summary Update a role
// @Description Update a role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body database.Role true "Role"
// @Success 200 {object} errorResponse "Role updated"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/role/{id} [put]
func (h *Handler) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid role ID")
		return
	}

	var role database.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.RoleRepository.UpdateRole(uint(id), role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated"})
}

// DeleteRole godoc
//
// @Summary Delete a role
// @Description Delete a role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} errorResponse "Role deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/role/{id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid role ID")
		return
	}

	err = h.Services.RoleRepository.DeleteRole(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
