package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// CreatePosition godoc
//
// @Summary Create a new position
// @Description Create a new position
// @Tags Position
// @Accept json
// @Produce json
// @Param position body database.Position true "Position"
// @Success 201 {object} errorResponse "Position created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/position [post]
func (h *Handler) CreatePosition(c *gin.Context) {
	var position database.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.PositionRepository.CreatePosition(position)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetPositionByID godoc
//
// @Summary Get a position
// @Description Get a position
// @Tags Position
// @Accept json
// @Produce json
// @Param id path int true "Position ID"
// @Success 200 {object} database.Position
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/position/{id} [get]
func (h *Handler) GetPositionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid position ID")
		return
	}

	position, err := h.Services.PositionRepository.GetPositionByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Position not found")
		return
	}

	c.JSON(http.StatusOK, position)
}

// UpdatePosition godoc
//
// @Summary Update a position
// @Description Update a position
// @Tags Position
// @Accept json
// @Produce json
// @Param id path int true "Position ID"
// @Param position body database.Position true "Position"
// @Success 200 {object} errorResponse "Position updated"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/position/{id} [put]
func (h *Handler) UpdatePosition(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid position ID")
		return
	}

	var position database.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.PositionRepository.UpdatePosition(uint(id), position)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Position updated"})
}

// DeletePosition godoc
//
// @Summary Delete a position
// @Description Delete a position
// @Tags Position
// @Accept json
// @Produce json
// @Param id path int true "Position ID"
// @Success 200 {object} errorResponse "Position deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/position/{id} [delete]
func (h *Handler) DeletePosition(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid position ID")
		return
	}

	err = h.Services.PositionRepository.DeletePosition(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Position deleted"})
}
