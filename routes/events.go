package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.api/models"
)

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

	event.UserID = ctx.GetInt64("userId")

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save data",
		})
		return
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

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid param",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update",
		})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
	})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid param",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		return
	}

	userId := ctx.GetInt64("userId")

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update",
		})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete operation failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
