package order

type Service interface {
}

type service struct {
	orderRepo Repository
}
