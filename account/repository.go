package account

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (rp *Repository) Login(username, password string) (string, error) {
	return username, nil
}
