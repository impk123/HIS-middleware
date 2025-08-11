package db

import (
	"fmt"
	"github.com/impk123/HIS-middleware/config"
	"github.com/impk123/HIS-middleware/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

// InitDB เชื่อมต่อฐานข้อมูลและเก็บ instance ไว้ใน package
func InitDB(cfg config.Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	fmt.Println(dsn)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อฐานข้อมูล: %v", err)
	}

	database = db
	log.Println("เชื่อมต่อฐานข้อมูลสำเร็จ")
}

func Migrate() {
	err := database.AutoMigrate(
		&models.Patient{}, // เพิ่ม model ที่ต้องการ migrate
		&models.Staff{},
	)
	if err != nil {
		log.Fatalf("Migration ล้มเหลว: %v", err)
	}
	log.Println("Migration สำเร็จ")

}

func GetDB() *gorm.DB {
	return database
}
