package main

import (
	"HIS-middleware/api"
	"HIS-middleware/config"
	"HIS-middleware/db"
	"github.com/gin-gonic/gin"
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
