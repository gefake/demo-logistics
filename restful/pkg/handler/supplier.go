package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// CreateSupplier godoc
// @Summary Create a new supplier
// @Description Create a new supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param supplier body database.Supplier true "Supplier object"
// @Success 201 {object} database.Supplier "Supplier created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/supplier [post]
func (h *Handler) CreateSupplier(c *gin.Context) {
	var supplier database.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.SupplierRepository.CreateSupplier(supplier)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetSupplierByID godoc
//
// @Summary Retrieve a supplier by ID
// @Description Retrieve a supplier by ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Param id path uint true "Supplier ID"
// @Success 200 {object} database.Supplier "Supplier retrieved"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /api/supplier/{id} [get]
func (h *Handler) GetSupplierByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid supplier ID")
		return
	}

	supplier, err := h.Services.SupplierRepository.GetSupplierByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Supplier not found")
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// UpdateSupplier godoc
//
// @Summary Update a supplier
// @Description Update a supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param id path uint true "Supplier ID"
// @Param supplier body database.Supplier true "Supplier object"
// @Success 200 {object} database.Supplier "Supplier updated"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/supplier/{id} [put]
func (h *Handler) UpdateSupplier(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid supplier ID")
		return
	}

	var supplier database.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.SupplierRepository.UpdateSupplier(uint(id), supplier)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated"})
}

// DeleteSupplier deletes a supplier by ID.
//
// @Summary Delete a supplier by ID
// @Description Delete a supplier by ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Param id path uint true "Supplier ID"
// @Success 200 {object} errorResponse "Supplier deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/supplier/{id} [delete]
func (h *Handler) DeleteSupplier(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid supplier ID")
		return
	}

	err = h.Services.SupplierRepository.DeleteSupplier(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted"})
}

// GetAllSuppliers godoc
//
// @Summary Get all suppliers
// @Description Get all suppliers
// @Tags Supplier
// @Accept json
// @Produce json
// @Success 200 {array} database.Supplier "Suppliers retrieved"
// @Failure 500 {object} errorResponse
// @Router /api/supplier [get]
func (h *Handler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.Services.SupplierRepository.GetAllSuppliers()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, suppliers)
}
