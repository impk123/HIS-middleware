package handlers

import (
	"HIS-middleware/db/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateStaff(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var staff models.Staff
		if err := c.ShouldBindJSON(&staff); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&staff).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, staff)
	}
}
