package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"size:50;uniqueIndex"`
	Password    string `gorm:"size:100"`
	Hospital    string `gorm:"size:50;not null"`
	FirstNameTh string `gorm:"size:100"`
	LastNameTh  string `gorm:"size:100"`
}
