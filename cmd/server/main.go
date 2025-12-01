package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proj/internal/handlers"
	"proj/internal/service"
	"proj/internal/storage"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	stor := storage.NewStorage()
	serv := service.NewService(stor)
	hand := handlers.NewHandler(serv)

	r := gin.Default()
	r.POST("/links", hand.Links)
	r.GET("/links", hand.GetAllLinks)
	r.POST("/links/report", hand.ReportLinks)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	sig := <-quit
	log.Printf("Received signal: %s. Shutting down...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited gracefully")
}
