package services

type Services struct {
	AlgosServiceI
}

type AlgosServiceI interface {
	FindSubstring(str string) string
}

func NewServices(service AlgosServiceI) *Services {
	return &Services{
		AlgosServiceI: service,
	}
}
