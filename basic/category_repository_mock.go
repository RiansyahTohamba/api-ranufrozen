package basic

import "github.com/stretchr/testify/mock"

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CategoryRepositoryMock) FindById(id string) *Category {
	args := repo.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(Category)
		return &category
	}
}
