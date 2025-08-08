package middleware

import (
	"HIS-middleware/api/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your authentication logic here
		// Example: Check JWT token or session
		c.Next()
	}
}
