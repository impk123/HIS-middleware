package migrations

import (
	"HIS-middleware/db/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Patient{},
		&models.Staff{},
	); err != nil {
		return err
	}
	return nil
}
