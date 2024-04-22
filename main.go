package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.api/db"
	"rest.api/models"
)

func main() {
	db.InitDB()
	
	server := gin.Default()

	server.GET("/event/:id", getEvent)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}

	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data provided",
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save data",
		})
	}

	ctx.JSON(http.StatusCreated, event)
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid param",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}

	ctx.JSON(http.StatusOK, event)

}