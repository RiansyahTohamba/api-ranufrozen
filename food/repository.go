package food

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	// FindById(id int) (Food, error)
	FindAll() ([]Food, error)
	Create(food Food) (Food, error)
	OptimisTx()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (fr *repository) OptimisTx() {
	statement := `SELECT id, val1, val2 FROM theTable WHERE iD = ?`
	// find by id bagaimana ya?
	row, err = fr.db.QueryRow(statement, id)

	// - {code that calculates new values}

	newVal1 := 10 + 10
	newVal2 := 10 + 30

	// oldVal1 := 10

	statupdate := `UPDATE theTable 
		SET val1 = @newVal1
		WHERE iD = @theId AND 
		val1 = @oldVal1;`

	tx, err := tr.db.Begin()

	_, err = tx.Exec(statupdate, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if AffectedRows == 1 {
		//    {go on with your other code}
		fmt.Println("Perubahan success")
	} else {
		tx.Rollback()
		//    {decide what to do since it has gone bad... in your code}
	}
}

// function milik struct 'repository'
// diawali (r *repository)
// func (r *repository) FindById(id int) (Food, error) {
// 	var food Food
// 	// hasil disimpan di adress food
// 	err := r.db.First(&food).Error
// 	return food, err
// }

func (r *repository) FindAll() ([]Food, error) {
	var foods []Food
	err := r.db.Find(&foods).Error
	return foods, err
}

func (r *repository) Create(food Food) (Food, error) {
	err := r.db.Create(&food).Error
	return food, err
}
