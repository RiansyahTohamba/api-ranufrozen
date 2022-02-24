package order

type Service interface {
	FindAll() ([]Order, error)
	// FindById(id int) (Order, error)
	// Create() (Order, error)
}

type service struct {
	orderRepo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]Order, error) {
	orders, err := s.orderRepo.FindAll()
	return orders, err
}

// func (s *service) Create() (Order, error) {
// 	orders, err := s.orderRepo.FindAll()
// 	Order
// 	return orders, err
// }
