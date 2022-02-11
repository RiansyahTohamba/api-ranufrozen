package food

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Food, error)
	FindByID(ID int) (Food, error)
	Create(food Food) (Food, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Food, error) {
	var foods []Food
	err := r.db.Find(&foods).Error
	return foods, err
}

func (r *repository) FindByID(ID int) (Food, error) {
	var food Food
	err := r.db.Find(&food, ID).Error
	return food, err
}

func (r *repository) Create(food Food) (Food, error) {
	err := r.db.Create(&food).Error
	return food, err
}
