package food

type Service interface {
}

// struct vs interface?
// type
type service struct {
	foodRepo Repository
}

// func (service Service) Get(id int) (*Food, error) {
// 	food := service.FoodRepo.FindById(id)
// 	if food == nil {
// 		return food, errors.New("Food Not Found")
// 	} else {
// 		return food, nil
// 	}

// }

// func GetMock(id int) Food {
// 	food := Food{
// 		name:  "Nugget",
// 		price: 20000,
// 	}
// 	return food

// }

// func (service Service) Create() {
// 	food := Food{}
// 	food.Name = "Nugget Ayam"
// 	food.Price = 25000.0
// 	food.Discount = 5
// 	food.Rating = 4
// 	food.Description = "Lezat sekali"

// 	foodSaved, err = foodRepo.Create(food)
// }
