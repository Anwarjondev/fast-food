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

// GetOrderByStatus returns a handler for getting orders by their status
// This is used for three different routes:
// - GET /orders/active
// - GET /orders/completed
// - GET /orders/all
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
