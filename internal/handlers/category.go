package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/gin-gonic/gin"
)

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get list of all food categories
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []repository.Category
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get category details by ID
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} repository.Category
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := repository.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}
