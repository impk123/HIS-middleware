package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/api/handlers"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		claims, err := handlers.ValidateJWTToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		// เซ็ตข้อมูลจาก token ลงใน context
		c.Set("staff_id", claims.StaffID)
		c.Set("hospital", claims.Hospital)
		c.Set("username", claims.Username)

		c.Next()
	}
}
