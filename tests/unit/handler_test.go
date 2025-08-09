package unit

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/api/handlers"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	return db
}

func TestCreateStaff(t *testing.T) {
	db := setupTestDB()
	handler := handlers.CreateStaff(db)

	// สร้าง test router
	r := gin.Default()
	r.POST("/staff/create", handler)

	// Test case 1: สร้าง staff สำเร็จ
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staff/create",
		bytes.NewBufferString(`{"username":"testuser","password":"testpass","hospital_id":1}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	// เพิ่ม test cases อื่นๆ ตามต้องการ
}
