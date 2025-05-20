package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/gin-gonic/gin"
)

// GetFoodsByCategory godoc
// @Summary Get foods by category
// @Description Get list of foods in a category
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} []repository.Food
// @Router /categories/{id}/foods [get]
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
