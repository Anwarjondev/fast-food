package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	Items []repository.OrderDetail `json:"items"`
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create a new food order
// @Tags orders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param order body CreateOrderInput true "Order details"
// @Success 200 {object} map[string]interface{}
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	userID := c.GetInt("user_id")
	var input CreateOrderInput

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderID, err := repository.CreateOrder(userID, input.Items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_id": orderID})
}

// GetOrderByStatus godoc
// @Summary Get orders by status
// @Description Get list of orders by status (active, completed, all)
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Param status query string false "Order status (active, completed, all)"
// @Success 200 {object} []repository.Order
// @Router /orders/active [get]
// @Router /orders/completed [get]
// @Router /orders/all [get]
func GetOrderByStatus(status string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")
		orders, err := repository.GetAllOrderByStatus(userID, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

// CancelOrder godoc
// @Summary Cancel order
// @Description Cancel an existing order
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Router /orders/{order_id} [put]
func CancelOrder(c *gin.Context) {
	userID := c.GetInt("user_id")
	orderID := c.Param("order_id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	err = repository.CancelOrder(userID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order canceled"})
}
