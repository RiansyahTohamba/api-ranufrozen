package food

type Service interface {
	// FindById(id int) (Food, error)
	// FindAll() ([]Food, error)
	Create(food FoodRequest) (Food, error)
}

// struct vs interface?
// type
type service struct {
	foodRepo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

// func (s *service) FindAll() ([]Food, error) {
// 	foods, err := s.foodRepo.FindAll()
// 	return foods, err
// }

// func (s *service) FindById(id int) (Food, error) {
// 	food, err := s.foodRepo.FindById(id)
// 	return food, err
// }

func (s *service) Create(foodReq FoodRequest) (Food, error) {
	price, _ := foodReq.Price.Float64()
	food := Food{
		Name:  foodReq.Name,
		Price: float64(price),
	}
	// disini harus ada pengondisian untuk price 'less or equal' dibawah zero
	newFood, err := s.foodRepo.Create(food)
	return newFood, err
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
