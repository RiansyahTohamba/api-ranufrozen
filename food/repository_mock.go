package food

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

// func (repo *RepositoryMock) FindById(id int) (Food, error) {
// arguments := repo.Mock.Called(id)
// if arguments.Get(0) == nil {
// 	return nil, nil
// } else {
// 	food := arguments.Get(0).(Food)
// 	return food, nil
// }
// 	food := Food{
// 		Name: "nugget",
// 	}

// 	return food, nil
// }

// func (repo *RepositoryMock) FindAll() ([]Food, error) {
// 	return nil, nil
// }

// func (repo *repository) Create(food Food) (Food, error) {
// 	nugget := Food{
// 		Name: "nugget",
// 	}
// 	return nugget, nil
// }
