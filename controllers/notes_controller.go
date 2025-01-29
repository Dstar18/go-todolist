package controllers

import (
	"go-todolist/database"
	"go-todolist/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetNotes(c *gin.Context) {
	var notes []models.Notes
	database.DB.Find(&notes)

	// return success
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "List Data User",
		"data":    notes,
	})
}

// store validaiton
type NotesVal struct {
	Title string `json:"title" validate:"required"`
}

var validateNotes = validator.New()

func StoreNotes(c *gin.Context) {
	// request struct validation
	var note NotesVal

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid request body",
		})
		return
	}

	// check vlaidation for v10
	if err := validateNotes.Struct(note); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": errors,
		})
		return
	}

	param := models.Notes{
		Title:       note.Title,
		IsCompleted: 0,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	// create to db
	if err := database.DB.Create(&param).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	// return success
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Created successfully",
		"data": map[string]interface{}{
			"id":    param.ID,
			"title": param.Title,
		},
	})
}

// store validaiton
type NotesValUpdate struct {
	Title       string `json:"title" validate:"required"`
	IsCompleted int    `json:"is_completed" validate:"numeric,oneof=0 1"` //status numeric 0-1
}

func UpdateNotes(c *gin.Context) {
	// request param id
	id := c.Param("id")

	// request struct
	var notesM models.Notes

	// check data by id
	if err := database.DB.First(&notesM, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Book Not Found",
			"data":    nil,
		})
		return
	}

	var noteParam NotesValUpdate

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&noteParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid request body",
		})
		return
	}

	// check vlaidation for v10
	if err := validateNotes.Struct(noteParam); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": errors,
		})
		return
	}

	// save to db
	notesM.Title = noteParam.Title
	notesM.IsCompleted = noteParam.IsCompleted
	notesM.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := database.DB.Save(&notesM).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	// return success
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Created successfully",
		"data":    notesM,
	})
}
