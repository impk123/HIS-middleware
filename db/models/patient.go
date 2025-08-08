package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	id             uint   `gorm:"primaryKey"`
	first_name_th  string `gorm:"size:100;not null"`
	middle_name_th string `gorm:"size:100"`
	last_name_th   string `gorm:"size:100;not null"`
	first_name_en  string `gorm:"size:100"`
	middle_name_en string `gorm:"size:100"`
	last_name_en   string `gorm:"size:100"`
	date_of_birth  string `gorm:"type:date"`
	patient_hn     string `gorm:"size:20;uniqueIndex"`
	national_id    string `gorm:"size:13;uniqueIndex"`
	passport_id    string `gorm:"size:20;uniqueIndex"`
	phone_number   string `gorm:"size:20"`
	email          string `gorm:"size:100"`
	gender         string `gorm:"size:1"` // M or F
	address        string `gorm:"size:254"`
	hosp_id        uint
}
