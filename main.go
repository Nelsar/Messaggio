package main

import (
	"github.com/gin-gonic/gin"
	"messaggio.com/handlers"
)

func main() {
	r := initHttp()
	r.Run(":8090")
}

func initHttp() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	route := gin.Default()
	route.POST("/createEvent", handlers.CreateEvent)
	route.PUT("/updateEvent", handlers.UpdateEvent)
	route.GET("/getevents", handlers.GetEvents)

	return route
}
