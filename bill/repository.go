package bill

import "gorm.io/gorm"

type Repository interface {
	FindById(id int) (Bill, error)
	FindAll() ([]Bill, error)
	// Create(bill Bill) (Bill, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindById(id int) (Bill, error) {
	var bill Bill
	err := repo.db.First(&bill).Error
	return bill, err
}
func (r *repository) FindAll() ([]Bill, error) {
	var bills []Bill
	err := r.db.Find(&bills).Error
	return bills, err
}
