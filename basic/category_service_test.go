package basic_test

import (
	"api-ranufrozen/basic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var catRepo = &basic.CategoryRepositoryMock{Mock: mock.Mock{}}
var catService = basic.CategoryService{Repository: catRepo}

func TestCategoryService_GetFound(t *testing.T) {
	category := basic.Category{
		Id:   "2",
		Name: "Hp",
	}
	catRepo.Mock.On("FindById", "2").Return(category)

	result, err := catService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}

func TestCategoryService_GetNotFound(t *testing.T) {
	catRepo.Mock.On("FindById", "1").Return(nil)
	category, err := catService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}
func TestCategoryService_GetSuccess(t *testing.T) {
	// category := basic.Category{
	// 	Id:   "10",
	// 	Name: "Laptop",
	// }
}
