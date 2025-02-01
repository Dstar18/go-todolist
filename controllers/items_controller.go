package controllers

import (
	"go-todolist/database"
	"go-todolist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowItems(c *gin.Context) {
	// request param id
	id := c.Param("id")

	// request struct
	var itemM models.Items

	// check data by id
	if err := database.DB.First(&itemM, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Item Not Found",
			"data":    nil,
		})
		return
	}

	// return success
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "show successfully",
		"data":    itemM,
	})
}

func DestroyItems(c *gin.Context) {
	// request param id
	id := c.Param("id")

	// request struct
	var itemM models.Items

	// check data by id
	if err := database.DB.First(&itemM, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Item Not Found",
			"data":    nil,
		})
		return
	}

	// delete to db
	if err := database.DB.Delete(&itemM).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete item",
			"error":   err.Error(),
		})
		return
	}

	// Respon berhasil
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Deleted successfully",
		"data": map[string]interface{}{
			"id": itemM.ID,
		},
	})
}
