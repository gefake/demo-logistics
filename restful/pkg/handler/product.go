package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body database.Product true "Product object"
// @Success 201 {object} database.Product "Product created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/product [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var product database.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.ProductRepository.CreateProduct(product)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path uint true "Product ID"
// @Success 200 {object} database.Product "Product retrieved"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /api/product/{id} [get]
func (h *Handler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.Services.ProductRepository.GetProductByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update product by ID
// @Description Update a product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path uint true "Product ID"
// @Param product body database.Product true "Product object"
// @Success 200 {object} database.Product "Product updated"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/product/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product database.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.ProductRepository.UpdateProduct(uint(id), product)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

// getCats godoc
// @Summary Получить все категории товаров
// @Description Получить все категории товаров
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} string "Массив категорий"
// @Failure 500 {object} errorResponse
// @Router /api/product/cats [get]
func (h *Handler) getCats(c *gin.Context) {
	cats, err := h.Services.ProductRepository.GetCategories()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cats)
}

// DeleteProduct godoc
// @Summary Удалить продукт по ID
// @Description Удаляет продукт по его идентификатору
// @Tags Product
// @Accept json
// @Produce json
// @Param id path uint true "ID продукта"
// @Success 200 {object} errorResponse "Product deleted"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /api/product/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.Services.ProductRepository.DeleteProduct(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// getAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} database.Product "Products retrieved"
// @Failure 500 {object} errorResponse
// @Router /api/product [get]
func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.Services.ProductRepository.GetAllProducts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
