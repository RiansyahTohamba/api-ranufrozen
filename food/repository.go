package food

import (
	"fmt"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(id int) Food
	FindAll() ([]Food, error)
	Create(food Food) (Food, error)
	BuyProduct(id int, quantity int)
}

type repository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewRepository(gorm *gorm.DB, rcl *redis.Client) *repository {
	return &repository{db: gorm, cache: rcl}
}

func (fr *repository) FindById(id int) Food {
	var food Food

	fr.db.Model(Food{ID: id}).First(&food)
	return food
}

func (r *repository) FindAll() ([]Food, error) {
	var foods []Food
	err := r.db.Find(&foods).Error
	return foods, err
}

func (r *repository) Create(food Food) (Food, error) {
	err := r.db.Create(&food).Error
	return food, err
}

// Optimis Transaction
func (fr *repository) BuyProduct(id int, quantity int) {

	fmt.Println("begin transaction !")

	fr.db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction
		// (use 'tx' from this point, not 'db')

		// bagaimana caranya GORM mapping struct Food dan table food?
		// padahal saya tidak buat secara explicit? struct Food = `table food`
		var food Food
		tx.Model(Food{ID: id}).First(&food)

		fmt.Println(food.Stock)

		newStock := food.Stock - quantity
		fmt.Printf("stok berkurang menjadi %d\n", newStock)

		dml := tx.Model(Food{}).Where("id = ?", id).Update("stock", newStock)

		if err := dml.Error; err != nil {
			// return any error will rollback
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	// - {code that calculates new values}

	// statupdate := `UPDATE theTable
	// 	SET val1 = @newVal1
	// 	WHERE iD = @theId AND
	// 	val1 = @oldVal1;`

	// tx, err := tr.db.Begin()

	// _, err = tx.Exec(statupdate, id)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// if AffectedRows == 1 {
	// 	//    {go on with your other code}
	// 	fmt.Println("Perubahan success")
	// } else {
	// 	//    {decide what to do since it has gone bad... in your code}
	// 	tx.Rollback()
	// }
}

// function milik struct 'repository'
// diawali (r *repository)
// func (r *repository) FindById(id int) (Food, error) {
// 	var food Food
// 	// hasil disimpan di adress food
// 	err := r.db.First(&food).Error
// 	return food, err
// }
