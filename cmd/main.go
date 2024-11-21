package main

import (
	"auth-service/config"
	"auth-service/internal/entities"
	"auth-service/internal/infrastructure/database"
	"auth-service/internal/interfaces/http"
	"auth-service/internal/usecases"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate database
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize dependencies
	userRepo := &database.PostgresUserRepository{DB: db}
	registerUseCase := &usecases.RegisterUserUseCase{UserRepo: userRepo}

	// Initialize handlers
	authHandler := &http.AuthHandler{RegisterUseCase: registerUseCase}

	// Initialize router
	router := gin.Default()

	router.POST("/auth/register", authHandler.Register)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Service is running"})
	})

	// Trust proxy
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Run server
	router.Run(":8080")
}
