package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	FirstNameTh  string `gorm:"size:100;not null"`
	MiddleNameTh string `gorm:"size:100"`
	LastNameTh   string `gorm:"size:100;not null"`
	FirstNameEn  string `gorm:"size:100"`
	MiddleNameEn string `gorm:"size:100"`
	LastNameEn   string `gorm:"size:100"`
	DateOfBirth  string `gorm:"type:date"`
	PatientHN    string `gorm:"size:20;uniqueIndex"`
	NationalID   string `gorm:"size:13;"`
	PassportID   string `gorm:"size:20;"`
	PhoneNumber  string `gorm:"size:20"`
	Email        string `gorm:"size:100"`
	Gender       string `gorm:"size:1"` // M or F
	Address      string `gorm:"size:254"`
	Hospital     string `gorm:"size:50;not null"`
}
