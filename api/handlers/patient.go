package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/impk123/HIS-middleware/db/models"
	"github.com/impk123/HIS-middleware/pkg/his"
	"gorm.io/gorm"
)

// CreatePatient - สร้างผู้ป่วย
func CreatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient models.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, patient)
	}
}

// SearchPatient - ค้นหาผู้ป่วย
func SearchPatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง hospital จาก staff ที่ล็อกอิน
		usernamex, exists := c.Get("username")
		staffHospital, exists := c.Get("hospital")
		fmt.Println(staffHospital)
		fmt.Print(usernamex)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// ดึง query parameters
		query := c.Request.URL.Query()
		searchParams := models.Patient{
			NationalID:  query.Get("national_id"),
			PassportID:  query.Get("passport_id"),
			FirstNameTh: query.Get("first_name"),
			LastNameTh:  query.Get("last_name"),
			PhoneNumber: query.Get("phone_number"),
			Email:       query.Get("email"),
			Hospital:    staffHospital.(string),
		}

		if nationalID := query.Get("national_id"); nationalID != "" {
			hospitalAClient := his.NewHospitalAClient("https://hospital-a.api.co.th")
			patient, err := hospitalAClient.SearchPatient(nationalID)
			if err == nil {
				// เก็บข้อมูลผู้ป่วยจาก Hospital A ลง database
				newPatient := models.Patient{
					FirstNameTh: patient.FirstNameTH,
					LastNameTh:  patient.LastNameTH,
					// ... map all fields ...
					Hospital: staffHospital.(string),
				}
				db.Create(&newPatient)
			}
		}

		// สร้าง dynamic query
		dbQuery := db.Model(&models.Patient{})
		if searchParams.NationalID != "" {
			dbQuery = dbQuery.Where("national_id = ?", searchParams.NationalID)
		}
		if searchParams.PassportID != "" {
			dbQuery = dbQuery.Where("passport_id = ?", searchParams.PassportID)
		}
		if searchParams.FirstNameTh != "" {
			dbQuery = dbQuery.Where("first_name_th LIKE ?", "%"+searchParams.FirstNameTh+"%")
		}
		if searchParams.LastNameTh != "" {
			dbQuery = dbQuery.Where("last_name_th LIKE ?", "%"+searchParams.LastNameTh+"%")
		}
		if searchParams.PhoneNumber != "" {
			dbQuery = dbQuery.Where("phone_number = ?", searchParams.PhoneNumber)
		}
		if searchParams.Email != "" {
			dbQuery = dbQuery.Where("email = ?", searchParams.Email)
		}
		dbQuery = dbQuery.Where("hospital = ?", searchParams.Hospital)

		var patients []models.Patient
		if err := dbQuery.Find(&patients).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// กรองข้อมูล sensitive ก่อนส่งกลับ
		var response []gin.H
		for _, p := range patients {
			response = append(response, gin.H{
				"first_name_th": p.FirstNameTh,
				"last_name_th":  p.LastNameTh,
				"first_name_en": p.FirstNameEn,
				"last_name_en":  p.LastNameEn,
				"date_of_birth": p.DateOfBirth,
				"patient_hn":    p.PatientHN,
				"phone_number":  p.PhoneNumber,
				"email":         p.Email,
				"gender":        p.Gender,
			})
		}

		c.JSON(http.StatusOK, response)
	}
}
