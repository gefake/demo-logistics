package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// CreateCargo Cargo godoc
// @Summary Create a new cargo
// @Description Create a new cargo
// @Tags cargo
// @Accept json
// @Produce json
// @Param data body database.Cargo true "Cargo data"
// @Success 201 {int} int "Cargo ID"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/cargo [post]
func (h *Handler) CreateCargo(c *gin.Context) {
	var cargo database.Cargo
	if err := c.ShouldBindJSON(&cargo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.CargoRepository.CreateCargo(cargo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetCargoByID Cargo godoc
// @Summary Get a cargo by ID
// @Description Get a cargo by ID
// @Tags cargo
// @Accept json
// @Produce json
// @Param id path uint true "Cargo ID"
// @Success 200 {object} database.Cargo
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/cargo/{id} [get]
func (h *Handler) GetCargoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid cargo ID")
		return
	}

	cargo, err := h.Services.CargoRepository.GetCargoByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Cargo not found")
		return
	}

	c.JSON(http.StatusOK, cargo)
}

// UpdateCargo Cargo godoc
// @Summary Update a cargo
// @Description Update a cargo
// @Tags cargo
// @Accept json
// @Produce json
// @Param id path uint true "Cargo ID"
// @Param data body database.Cargo true "Cargo data"
// @Success 200 {string} string "Cargo updated"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/cargo/{id} [put]
func (h *Handler) UpdateCargo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid cargo ID")
		return
	}

	var cargo database.Cargo
	if err := c.ShouldBindJSON(&cargo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.CargoRepository.UpdateCargo(uint(id), cargo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cargo updated"})
}

// DeleteCargo Cargo godoc
// @Summary Delete a cargo
// @Description Delete a cargo
// @Tags cargo
// @Accept json
// @Produce json
// @Param id path uint true "Cargo ID"
// @Success 200 {string} string "Cargo deleted"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/cargo/{id} [delete]
func (h *Handler) DeleteCargo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid cargo ID")
		return
	}

	err = h.Services.CargoRepository.DeleteCargo(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cargo deleted"})
}

type successResponseAllCargos struct {
	Сargos    []database.Cargo `json:"cargos"`
	PageCount int              `json:"page_count"`
}

// GetAllCargos Cargo godoc
// @Summary Get all cargos
// @Description Get all cargos with pagination
// @Tags cargo
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param perPage query int false "Number of cargos per page"
// @Success 200 {object} successResponseAllCargos
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/cargos [get]
func (h *Handler) GetAllCargos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	perPage, err := strconv.Atoi(c.Query("perPage"))
	if err != nil || perPage <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "Invalid perPage value")
		return
	}

	cargos, totalCount, err := h.Services.CargoRepository.GetAllCargos(page, perPage)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	test := successResponseAllCargos{cargos, totalCount}

	c.JSON(http.StatusOK, test)
}

// GetCargosByUserID godoc
//
// @Summary Get all cargos by user ID
// @Description Get all cargos by user ID
// @Tags cargo
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} database.Cargo
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/cargo/user/{id} [get]
func (h *Handler) GetCargosByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	cargos, err := h.Services.CargoRepository.GetCargoOrdersByUserID(uint(userID))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cargos)
}
