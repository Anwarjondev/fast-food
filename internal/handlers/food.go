package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetFoodsByCategory(c *gin.Context) {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	foods, err := repository.GetFoodsByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}
