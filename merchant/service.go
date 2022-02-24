package merchant

type Service interface {
}

type service struct {
	merchRepo Repository
}

func NewService(merchRepo Repository) *service {
	return &service{merchRepo}
}
