package main

import (
	"fmt"
	"github.com/yourusername/url-shortener/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/yourusername/url-shortener/internal/handler"
	"github.com/yourusername/url-shortener/internal/storage"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := storage.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	sv := service.NewGenerator(db)
	re := service.NewRedirect(db)
	ana := service.NewAnalytics(db)
	h := handler.NewHandler(sv, re, ana)

	r := gin.Default()

	r.POST("/api/shorten", h.CreateShortURL)
	r.GET("/api/analytics/:code", h.GetAnalytics)
	r.GET(":code", h.Redirect)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("listening on %s", addr)
	r.Run(addr)
}
