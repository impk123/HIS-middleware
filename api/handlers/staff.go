package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateStaff(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required" example:"admin"`
			Password string `json:"password" binding:"required" example:"admin123"`
			Hospital string `json:"hospital" binding:"required" example:"12123"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// เข้ารหัส password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}

		staff := models.Staff{
			Username: request.Username,
			Password: string(hashedPassword),
			Hospital: request.Hospital,
		}

		if err := db.Create(&staff).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username already exists"})
			return
		}

		// ไม่ส่ง password กลับ
		staff.Password = ""
		c.JSON(http.StatusCreated, staff)
	}
}

func StaffLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required" example:"admin"`
			Password string `json:"password" binding:"required" example:"admin123"`
			Hospital string `json:"hospital" binding:"required" example:"12123"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var staff models.Staff
		if err := db.Where("username = ? AND hospital = ?", request.Username, request.Hospital).First(&staff).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(request.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// สร้าง JWT token
		token, err := GenerateJWTToken(staff.ID, staff.Hospital, staff.Username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		fmt.Println(token)
		// ส่ง token กลับไปให้ client
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"staff": gin.H{
				"id":       staff.ID,
				"username": staff.Username,
				"hospital": staff.Hospital,
			},
			"expires_in": 86400, // 24 ชั่วโมงในวินาที
			"token_type": "Bearer",
		})
	}
}

// func generateJWTToken(staffID uint, Hospital string) (string, error) {
// 	// Implement JWT generation here
// 	// ตัวอย่างใช้ library เช่น github.com/dgrijalva/jwt-go
// 	return "", nil
// }
