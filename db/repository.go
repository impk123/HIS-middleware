package db

import (
	"github.com/impk123/HIS-middleware/db/models"

	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	GetAllPatients() ([]models.Patient, error)
	GetPatientByID(id uint) (models.Patient, error)
	CreatePatient(patient models.Patient) error
	UpdatePatient(patient models.Patient) error
	DeletePatient(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	result := r.db.Find(&patients)
	return patients, result.Error
}

func (r *repository) GetPatientByID(id uint) (models.Patient, error) {
	var patient models.Patient
	result := r.db.First(&patient, id)
	return patient, result.Error
}

func (r *repository) CreatePatient(patient models.Patient) error {
	return r.db.Create(&patient).Error
}

func (r *repository) UpdatePatient(patient models.Patient) error {
	return r.db.Save(&patient).Error
}

func (r *repository) DeletePatient(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}
