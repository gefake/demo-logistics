package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

//func (h *Handler) getWarehouses(c *gin.Context) {
//	var warehouses []database.Warehouse
//	h.Services.Repo
//	c.JSON(http.StatusOK, gin.H{"data": warehouses})
//}

// createWarehouse godoc
// @Summary Create a new warehouse
// @Description Create a new warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param warehouse body database.Warehouse true "Warehouse object"
// @Success 201 {int} int "Warehouse ID"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/warehouse [post]
func (h *Handler) createWarehouse(c *gin.Context) {
	var input database.Warehouse

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	warehouseID, err := h.Services.WarehouseRepository.CreateWarehouse(input)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": warehouseID})
}

// updateWarehouse godoc
// @Summary Update an existing warehouse
// @Description Update an existing warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Param warehouse body database.Warehouse true "Warehouse object"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/warehouse/{id} [patch]
func (h *Handler) updateWarehouse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input database.Warehouse

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.WarehouseRepository.UpdateWarehouse(uint(id), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// deleteWarehouse godoc
// @Summary Delete an existing warehouse
// @Description Delete an existing warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/warehouse/{id} [delete]
func (h *Handler) deleteWarehouse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.Services.WarehouseRepository.DeleteWarehouse(uint(id))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// getWarehouseByID godoc
// @Summary Get an existing warehouse
// @Description Get an existing warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} database.Warehouse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/warehouse/{id} [get]
func (h *Handler) getWarehouseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	warehouse, err := h.Services.WarehouseRepository.GetWarehouseByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, warehouse)
}

// getAllWarehouses godoc
// @Summary Get all warehouses
// @Description Get all warehouses
// @Tags Warehouse
// @Accept json
// @Produce json
// @Success 200 {array} database.Warehouse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/warehouse [get]
func (h *Handler) getAllWarehouses(c *gin.Context) {
	warehouses, err := h.Services.WarehouseRepository.GetAllWarehouses()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, warehouses)
}
