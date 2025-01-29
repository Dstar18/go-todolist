package controllers

import (
	"go-todolist/database"
	"go-todolist/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	// return success
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "List Data User",
		"data":    users,
	})
}

// struct validation
type StoreInput struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=30"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=30"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=20"`
}

var validate = validator.New()

func StoreUser(c *gin.Context) {

	// request struct validation
	var user StoreInput

	// Request Post Paramaeter, and check type parameter
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// check vlaidation for v10
	if err := validate.Struct(user); err != nil {
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

	// request struct model
	var userM models.User

	checkEmail := database.DB.Where("email = ?", user.Email).First(&userM)
	if checkEmail.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": user.Email + " is already registered.",
		})
		return
	}

	param := models.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
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
			"email": param.Email,
		},
	})

}
