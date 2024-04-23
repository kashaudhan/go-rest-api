package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.api/models"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data provided",
		})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H {
		"message": "User create successfully",
	})

}