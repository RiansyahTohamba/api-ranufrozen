package basic

type CategoryRepository interface {
	FindById(id string) *Category
}
