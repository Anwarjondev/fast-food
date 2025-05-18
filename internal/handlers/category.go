package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

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
