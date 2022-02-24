package order

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindAll() ([]Order, error) {
	var orders []Order
	err := repo.db.Find(&orders).Error
	return orders, err
}
