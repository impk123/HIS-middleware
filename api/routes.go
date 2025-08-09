package api

import (
	"github.com/impk123/HIS-middleware/api/handlers"
	"github.com/impk123/HIS-middleware/api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Staff APIs
	r.POST("/staff/create", handlers.CreateStaff(db))
	r.POST("/staff/login", handlers.StaffLogin(db))

	// Patient APIs (protected)
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/patient/create", handlers.CreatePatient(db))
		authGroup.GET("/patient/search", handlers.SearchPatient(db))
	}
}
