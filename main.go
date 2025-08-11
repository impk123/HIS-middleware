package main

import (
	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/api"
	"github.com/impk123/HIS-middleware/config"
	"github.com/impk123/HIS-middleware/db"

	"fmt"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db.InitDB(*cfg)

	// migrate ตาราง
	fmt.Println("migrate ตาราง")
	db.Migrate()

	// // Initialize test database
	// db.InitDBTest(*cfg)

	// // migrate ตารางทดสอบ
	// fmt.Println("migrate ตารางทดสอบ")
	// db.MigrateTest()

	// Set up Gin router
	r := gin.Default()

	// Set up routes
	api.SetupRoutes(r, db.GetDB())

	// Start server
	r.Run(":8080")
}
