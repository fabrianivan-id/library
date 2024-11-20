package main

import (
	"fmt"
	"library-backend/handlers"
	"library-backend/middleware"
	"library-backend/models"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	AppName   string                `yaml:"app_name"`
	Port      string                `yaml:"port"`
	JwtSecret string                `yaml:"jwt_secret"`
	Database  models.DatabaseConfig `yaml:"database"`
}

func loadConfig() (*AppConfig, error) {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg AppConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	// Load configuration
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Connect to database
	err = models.ConnectDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize Gin
	r := gin.Default()

	// Public routes
	r.POST("/login", handlers.Login)
	r.GET("/docs", handlers.GetDocumentation)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg.JwtSecret))
	{
		auth.GET("/books", handlers.GetBooks)
		auth.POST("/books", middleware.AdminMiddleware(), handlers.AddBook)
		auth.PUT("/books/:id", middleware.AdminMiddleware(), handlers.UpdateBookStock)

		auth.POST("/transactions", handlers.BorrowBook)
		auth.PUT("/transactions/:id/return", handlers.ReturnBook)
		auth.GET("/transactions", middleware.AdminMiddleware(), handlers.GetTransactions)
	}

	// Start the server
	log.Printf("Starting %s on port %s", cfg.AppName, cfg.Port)
	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
