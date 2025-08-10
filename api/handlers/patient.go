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
		staffHospital, exists := c.Get("hospital")
		fmt.Println(staffHospital)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "login unauthorized"})
			return
		}

		// รับค่า body json จาก postman
		var patient models.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if nationalID := patient.NationalID; nationalID != "" {
			hospitalAClient := his.NewHospitalAClient("https://hospital-a.api.co.th")
			patient, err := hospitalAClient.SearchPatient(nationalID)
			if err == nil {
				// เก็บข้อมูลผู้ป่วยจาก Hospital A ลง database
				newPatient := models.Patient{
					FirstNameTh: patient.FirstNameTH,
					LastNameTh:  patient.LastNameTH,
					NationalID:  nationalID,
					// ... map all fields ...
					Hospital: staffHospital.(string),
				}
				db.Create(&newPatient)
			}
		}

		// สร้าง dynamic query
		dbQuery := db.Model(&models.Patient{})
		if patient.NationalID != "" {
			dbQuery = dbQuery.Where("national_id = ? AND hospital = ?", patient.NationalID, staffHospital)
		}
		if patient.PassportID != "" {
			dbQuery = dbQuery.Where("passport_id = ? AND hospital = ?", patient.PassportID, staffHospital)
		}
		if patient.FirstNameTh != "" {
			dbQuery = dbQuery.Where("first_name_th LIKE ? AND hospital = ?", "%"+patient.FirstNameTh+"%", staffHospital)
			fmt.Println(dbQuery)
		}
		if patient.LastNameTh != "" {
			dbQuery = dbQuery.Where("last_name_th LIKE ? AND hospital = ?", "%"+patient.LastNameTh+"%", staffHospital)
		}
		if patient.PhoneNumber != "" {
			dbQuery = dbQuery.Where("phone_number = ? AND hospital = ?", patient.PhoneNumber, staffHospital)
		}
		if patient.Email != "" {
			dbQuery = dbQuery.Where("email = ? AND hospital = ?", patient.Email, staffHospital)
		}
		// ถ้าไม่ได้รับ parameter จาก query ให้ดึงทั้งหมด
		if patient.NationalID == "" && patient.PassportID == "" && patient.FirstNameTh == "" && patient.LastNameTh == "" && patient.PhoneNumber == "" && patient.Email == "" {
			dbQuery = dbQuery.Where("hospital = ?", staffHospital)
		}

		var patients []models.Patient
		if err := dbQuery.Find(&patients).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		if len(patients) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		// // กรองข้อมูล sensitive ก่อนส่งกลับ
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
