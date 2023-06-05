package hotel

import (
	"gorm.io/gorm"
)

// public
type Repository interface {
	Save(hotel Hotel) (Hotel, error)
	FindAll() ([]Hotel, error)
	Destroy() error
}

// private
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(hotel Hotel) (Hotel, error) {
	err := r.db.Create(&hotel).Error
	if err != nil {
		return hotel, err
	}

	return hotel, nil
}

func (r *repository) FindAll() ([]Hotel, error) {
	var hotel []Hotel
	err := r.db.Find(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (r *repository) Destroy() error {
	err := r.db.Exec("TRUNCATE TABLE hotel").Error
	if err != nil {
		return err
	}
	return nil
}
