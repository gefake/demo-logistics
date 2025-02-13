package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/database"
	"net/http"
	"strconv"
	"time"
)

// CreateOrder godoc
//
// @Summary Create a new order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body database.Order true "Order"
// @Success 201 {object} errorResponse "Order created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var order database.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	order.OrderDate = time.Now()

	id, err := h.Services.OrderRepository.CreateOrder(order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetOrderByID godoc
//
// @Summary Get an order
// @Description Get an order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} database.Order
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order/{id} [get]
func (h *Handler) GetOrderByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := h.Services.OrderRepository.GetOrderByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Order not found")
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder godoc
//
// @Summary Update an order
// @Description Update an order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body database.Order true "Order"
// @Success 200 {object} errorResponse "Order updated"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var order database.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.OrderRepository.UpdateOrder(uint(id), order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}

// DeleteOrder godoc
//
// @Summary Delete an order
// @Description Delete an order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} errorResponse "Order deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order/{id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	err = h.Services.OrderRepository.DeleteOrder(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

// GetAllOrders godoc
//
// @Summary Get all orders
// @Description Get all orders
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} database.Order
// @Failure 500 {object} errorResponse
// @Router /api/order [get]
func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.Services.OrderRepository.GetAllOrders()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

// ORDER PRODUCTS

// CreateOrderProduct godoc
//
// @Summary Create a new order product
// @Description Create a new order product
// @Tags Order
// @Accept json
// @Produce json
// @Param order body database.OrderItem true "Order"
// @Success 201 {object} errorResponse "Order created"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order/products [post]
func (h *Handler) CreateOrderProduct(c *gin.Context) {
	var order database.OrderItem

	if err := c.ShouldBindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := order.CreateOrderItem(order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// DeleteOrderProduct godoc
//
// @Summary Delete an order product
// @Description Delete an order product
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} errorResponse "Order deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order/products/{id} [delete]
func (h *Handler) DeleteOrderProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	itm := database.OrderItem{}

	err = itm.DeleteOrderItem(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
