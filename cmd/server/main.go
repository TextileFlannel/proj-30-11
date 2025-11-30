package main

import (
	"proj/internal/handlers"
	"proj/internal/service"
	"proj/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	stor := storage.NewStorage()
	serv := service.NewService(stor)
	hand := handlers.NewHandler(serv)

	r := gin.Default()
	r.POST("/links", hand.Links)
	r.GET("/getAllLinks", hand.GetAllLinks)
	r.POST("/reportLinks", hand.ReportLinks)

	r.Run()
}
