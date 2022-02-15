package food

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (repo *RepositoryMock) FindById(id int) *Food {
	arguments := repo.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		food := arguments.Get(0).(Food)
		return &food
	}
}
