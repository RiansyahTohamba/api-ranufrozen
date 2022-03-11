package basic

import "errors"

type CategoryService struct {
	Repository CategoryRepository
}

func (catSer CategoryService) Get(id string) (*Category, error) {
	category := catSer.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category is Not Found")
	} else {
		return category, nil
	}
}
