package bill

type Service interface {
	FindAll() ([]Bill, error)
}

type service struct {
	billRepo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (serv *service) FindAll() ([]Bill, error) {
	bills, error := serv.billRepo.FindAll()
	return bills, error
}
