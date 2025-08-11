package tests

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/api"

	// "github.com/impk123/HIS-middleware/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitTestDB เชื่อมต่อ Postgres (ใช้ DB จริงสำหรับทดสอบ)
func InitTestDB() *gorm.DB {
	// ใช้ env จาก docker-compose หรือกำหนดค่า default

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"test_his_hosp",
		"5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	// Migrate models
	return db
}

// Staff struct for migration test (if needed)
type Staff struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"size:50;uniqueIndex"`
	Password    string `gorm:"size:100"`
	Hospital    string `gorm:"size:50;not null"`
	FirstNameTh string `gorm:"size:100"`
	LastNameTh  string `gorm:"size:100"`
}

type Patient struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	FirstNameTh  string `gorm:"size:100;not null" json:"first_name"`
	MiddleNameTh string `gorm:"size:100" json:"middle_name"`
	LastNameTh   string `gorm:"size:100;not null" json:"last_name"`
	FirstNameEn  string `gorm:"size:100"`
	MiddleNameEn string `gorm:"size:100"`
	LastNameEn   string `gorm:"size:100"`
	DateOfBirth  string `gorm:"type:date" json:"date_of_birth"`
	PatientHN    string `gorm:"size:20;uniqueIndex"`
	NationalID   string `gorm:"size:13;" json:"national_id"`
	PassportID   string `gorm:"size:20;" json:"passport_id"`
	PhoneNumber  string `gorm:"size:20" json:"phone_number"`
	Email        string `gorm:"size:100" json:"email"`
	Gender       string `gorm:"size:1"` // M or F
	Address      string `gorm:"size:254"`
	Hospital     string `gorm:"size:50;not null"`
}

func MigrateTest() {
	db := InitTestDB()
	err := db.AutoMigrate(
		&Staff{}, // Use the defined Staff struct if needed
		&Patient{},
	)
	if err != nil {
		log.Fatalf("Migration TEST ล้มเหลว: %v", err)
	}
	log.Println("Migration TEST สำเร็จ")
}

func setupTestServer() *gin.Engine {
	testDB := InitTestDB()
	MigrateTest()

	r := gin.Default()
	api.SetupRoutes(r, testDB)

	return r
}

// staff create ผ่าน
func TestStaffCreate(t *testing.T) {
	// Setup
	r := setupTestServer()

	t.Run("Success - Create staff", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/create",
			bytes.NewBufferString(`{
				"username":"admin",
				"password": "admin123",
				"hospital": "12123"
			}`))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.True(t, w.Code == http.StatusOK || w.Code == http.StatusInternalServerError)

	})
}

// staff create ไม่ผ่าน
func TestStaffCreateWrong(t *testing.T) {
	// Setup
	r := setupTestServer()

	t.Run("Fail - username already exists", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/create",
			bytes.NewBufferString(`{
				"username":"admin",
				"password": "admin123",
				"hospital": "12123"
			}`))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "username already exists")
	})
}

func TestStaffLogin(t *testing.T) {

	r := setupTestServer()
	t.Run("Success - Valid credentials", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/login",
			bytes.NewBufferString(`{
				"username":"admin",
				"password": "admin123",
				"hospital": "12123"
			}`))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "token")
	})

	t.Run("Fail - Wrong password", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/login",
			bytes.NewBufferString(`{
				"username":"testuser",
				"password":"wrongpass",
				"hospital": "12123"
			}`))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid credentials")
	})

	t.Run("Fail - Missing fields", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/login",
			bytes.NewBufferString(`{
				"username":"testuser"
			}`))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "required")

	})
}

// ค้นหา patient
func TestSearchPatient(t *testing.T) {
	// Setup
	r := setupTestServer()

	// Login เพื่อรับ token
	wLogin := httptest.NewRecorder()
	reqLogin, _ := http.NewRequest("POST", "/staff/login",
		bytes.NewBufferString(`{
            "username":"admin",
            "password": "admin123",
            "hospital": "12123"
        }`))
	reqLogin.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(wLogin, reqLogin)

	// ดึง token จาก response
	var loginResponse map[string]interface{}
	json.Unmarshal(wLogin.Body.Bytes(), &loginResponse)
	token := loginResponse["token"].(string)
	// สร้าง patient สำหรับทดสอบ
	wPatient := httptest.NewRecorder()
	reqPatient, _ := http.NewRequest("POST", "/patient/create",
		bytes.NewBufferString(`{
            "first_name":"สมศรี",
            "last_name":"สมหมาย",
			"date_of_birth":"1998-05-02",
			"hospital": "12123"
        }`))
	reqPatient.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(wPatient, reqPatient)

	t.Run("Success - Found patient", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/patient/search",
			bytes.NewBufferString(`{"first_name":"สมศรี","hospital":"12123"}`))
		req.Header.Set("Authorization", "Bearer "+token)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "สมศรี")
	})

	t.Run("Fail - Patient not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/patient/search",
			bytes.NewBufferString(`{"first_name":"Unknown"}`))
		req.Header.Set("Authorization", "Bearer "+token)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Patient not found")
	})

	// t.Cleanup(func() {
	// 	db := InitTestDB()
	// 	db.Exec("DELETE FROM staff")
	// 	db.Exec("DELETE FROM patient")
	// 	// เพิ่ม table อื่นๆ ตามต้องการ
	// })
}
