package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// CreateDeliveryRoute godoc
//
// @Summary Create a new delivery route
// @Description Create a new delivery route
// @Tags Delivery
// @Accept json
// @Produce json
// @Param delivery_route body database.DeliveryRoute true "Delivery route"
// @Success 201 {object} errorResponse "Delivery route created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery [post]
func (h *Handler) CreateDeliveryRoute(c *gin.Context) {
	var deliveryRoute database.DeliveryRoute
	if err := c.ShouldBindJSON(&deliveryRoute); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.DeliveryRouteRepository.CreateDeliveryRoute(deliveryRoute)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetDeliveryRouteByID godoc
//
// @Summary Get delivery route by ID
// @Description Get delivery route by ID
// @Tags Delivery
// @Accept json
// @Produce json
// @Param id path int true "Delivery route ID"
// @Success 200 {object} database.DeliveryRoute
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery/{id} [get]
func (h *Handler) GetDeliveryRouteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid delivery route ID")
		return
	}

	deliveryRoute, err := h.Services.DeliveryRouteRepository.GetDeliveryRouteByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Delivery route not found")
		return
	}

	c.JSON(http.StatusOK, deliveryRoute)
}

// UpdateDeliveryRoute godoc
//
// @Summary Update delivery route
// @Description Update delivery route
// @Tags Delivery
// @Accept json
// @Produce json
// @Param id path int true "Delivery route ID"
// @Param delivery_route body database.DeliveryRoute true "Delivery route"
// @Success 200 {object} errorResponse "Delivery route updated"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery/{id} [put]
func (h *Handler) UpdateDeliveryRoute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid delivery route ID")
		return
	}

	var deliveryRoute database.DeliveryRoute
	if err := c.ShouldBindJSON(&deliveryRoute); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.DeliveryRouteRepository.UpdateDeliveryRoute(uint(id), deliveryRoute)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery route updated"})
}

// DeleteDeliveryRoute godoc
//
// @Summary Delete delivery route
// @Description Delete delivery route
// @Tags Delivery
// @Accept json
// @Produce json
// @Param id path int true "Delivery route ID"
// @Success 200 {object} errorResponse "Delivery route deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/delivery/{id} [delete]
func (h *Handler) DeleteDeliveryRoute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid delivery route ID")
		return
	}

	err = h.Services.DeliveryRouteRepository.DeleteDeliveryRoute(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery route deleted"})
}

// GetAllRoutes godoc
//
// @Summary Get all routes
// @Description Get all routes
// @Tags Delivery
// @Accept json
// @Produce json
// @Success 200 {array} database.DeliveryRoute
// @Failure 500 {object} errorResponse
// @Router /api/delivery [get]
func (h *Handler) GetAllRoutes(c *gin.Context) {
	routes, err := h.Services.DeliveryRouteRepository.GetAllRoutes()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, routes)
}
