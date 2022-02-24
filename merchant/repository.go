package merchant

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Merchant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindAll() ([]Merchant, error) {
	var merchants []Merchant
	err := repo.db.Find(&merchants).Error
	return merchants, err
}
