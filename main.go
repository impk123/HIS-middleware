package main

import (
	"github.com/gin-gonic/gin"
	"hospital-middleware/api"
	"hospital-middleware/config"
	"hospital-middleware/db"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db.InitDB(cfg)

	// Set up Gin router
	r := gin.Default()

	// Set up routes
	api.SetupRoutes(r, db.GetDB())

	// Start server
	r.Run(":8080")
}
