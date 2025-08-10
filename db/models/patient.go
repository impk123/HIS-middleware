package models

import "gorm.io/gorm"

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
