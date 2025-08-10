package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	// "github.com/impk123/HIS-middleware/models"
	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/api"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitTestDB เชื่อมต่อ Postgres (ใช้ DB จริงสำหรับทดสอบ)
func InitTestDB() *gorm.DB {
	// ใช้ env จาก docker-compose หรือกำหนดค่า default
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "db" // ชื่อ service ใน docker-compose
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "HIS_hosp"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	// Migrate models
	// db.AutoMigrate(&models.Staff{}, &models.Patient{})

	return db
}

func setupTestServer() *gin.Engine {
	testDB := InitTestDB()

	r := gin.Default()
	api.SetupRoutes(r, testDB)

	return r
}

func TestFullAuthFlow(t *testing.T) {
	r := setupTestServer()

	// 1. ทดสอบล็อกอิน
	t.Run("Login to get token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/login",
			bytes.NewBufferString(`{
				"username":"admin",
				"password":"admin123",
				"hospital":"12123"
			}`))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "token")
	})

	// 2. ทดสอบใช้ token เพื่อเข้าถึง API ที่ต้องล็อกอิน
	t.Run("Access protected route with token", func(t *testing.T) {
		// ล็อกอินเพื่อรับ token ก่อน
		wLogin := httptest.NewRecorder()
		reqLogin, _ := http.NewRequest("POST", "/staff/login",
			bytes.NewBufferString(`{
				"username":"admin",
				"password":"admin123",
				"hospital":"12123"
			}`))
		reqLogin.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(wLogin, reqLogin)

		var loginResponse map[string]interface{}
		json.Unmarshal(wLogin.Body.Bytes(), &loginResponse)
		token := loginResponse["token"].(string)

		// ใช้ token เพื่อเข้าถึง API ที่ต้องล็อกอิน
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/patient/search", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
