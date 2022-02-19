package bill

type Service interface {
}

type service struct {
	billRepo Repository
}
