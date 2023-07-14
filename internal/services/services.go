package services

import "github.com/Levap123/tsarka-test-tasks/internal/services/algos"

type Services struct {
	AlgosServiceI
}

type AlgosServiceI interface {
	FindSubstring(str string) string
	EmailCheck(str string) []string
}

func NewServices() *Services {
	return &Services{
		AlgosServiceI: algos.NewService(),
	}
}
