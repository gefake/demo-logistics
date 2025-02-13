package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// getDeliverySchedules godoc
// @Summary Get all delivery schedules
// @Description Get all delivery schedules
// @Tags DeliverySchedule
// @Accept json
// @Produce json
// @Success 200 {array} database.DeliverySchedule
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery-schedule [get]
func (h *Handler) getAllDeliverySchedules(c *gin.Context) {
	schedules, err := h.Services.DeliveryScheduleRepository.GetAllDeliverySchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedules)
}

// getDeliveryScheduleByID godoc
// @Summary Get an existing delivery schedule
// @Description Get an existing delivery schedule
// @Tags DeliverySchedule
// @Accept json
// @Produce json
// @Param id path int true "DeliverySchedule ID"
// @Success 200 {object} database.DeliverySchedule
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery-schedule/{id} [get]
func (h *Handler) getDeliveryScheduleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	schedule, err := h.Services.DeliveryScheduleRepository.GetDeliveryScheduleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

// createDeliverySchedule godoc
// @Summary Create a new delivery schedule
// @Description Create a new delivery schedule
// @Tags DeliverySchedule
// @Accept json
// @Produce json
// @Param delivery-schedule body database.DeliverySchedule true "DeliverySchedule object"
// @Success 201 {string} string "ok"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery-schedule [post]
func (h *Handler) createDeliverySchedule(c *gin.Context) {
	var input database.DeliverySchedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduleID, err := h.Services.DeliveryScheduleRepository.CreateDeliverySchedule(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": scheduleID})
}

// updateDeliverySchedule godoc
// @Summary Update an existing delivery schedule
// @Description Update an existing delivery schedule
// @Tags DeliverySchedule
// @Accept json
// @Produce json
// @Param id path int true "DeliverySchedule ID"
// @Param delivery-schedule body database.DeliverySchedule true "DeliverySchedule object"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery-schedule/{id} [patch]
func (h *Handler) updateDeliverySchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	var input database.DeliverySchedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Services.DeliveryScheduleRepository.UpdateDeliverySchedule(uint(id), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// deleteDeliverySchedule godoc
// @Summary Delete an existing delivery schedule
// @Description Delete an existing delivery schedule
// @Tags DeliverySchedule
// @Accept json
// @Produce json
// @Param id path int true "DeliverySchedule ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery-schedule/{id} [delete]
func (h *Handler) deleteDeliverySchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	err = h.Services.DeliveryScheduleRepository.DeleteDeliverySchedule(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
