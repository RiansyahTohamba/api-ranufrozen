# struct?
kalau struct sudah concrete, misalkan pada file repository.go
type repository struct {
	db *gorm.DB
}


# interface?
kalau interface masih abstract, tinggal di extend lagi. misalkan pada file repository.go

type Repository interface {
	FindAll() ([]Food, error)
	FindByID(ID int) (Food, error)
	Create(food Food) (Food, error)
}
