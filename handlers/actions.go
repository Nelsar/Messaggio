package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"messaggio.com/amqp"
	"messaggio.com/db"
	"messaggio.com/models"
)

func CreateEvent(ctx *gin.Context) {
	event := models.Event{}
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.AddEvent(event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = amqp.ProducerHanler(ctx, event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UpdateEvent(ctx *gin.Context) {
	event := models.Event{}
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.UpdateEvent(event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetEvents(ctx *gin.Context) {
	events, err := db.GetEvents()

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"events": events})
}
