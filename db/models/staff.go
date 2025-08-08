package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	id            uint   `gorm:"primaryKey"`
	Username      string `gorm:"size:50;uniqueIndex"`
	Password      string `gorm:"size:100"`
	hosp_id       uint
	first_name_th string `gorm:"size:100"`
	last_name_th  string `gorm:"size:100"`
}
