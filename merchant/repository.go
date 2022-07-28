package merchant

// standard library go
//
// third-party go
import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindAll() ([]Merchant, error) {
	var merchants []Merchant
	err := repo.db.Find(&merchants).Error
	return merchants, err
}
