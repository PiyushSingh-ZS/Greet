package name

import (
	helloSvc "awesomeProject/internal/service/hello"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

var hellos = helloSvc.New()

func (s *Service) Greet(name string) (string, error) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	return hellos.Greet(ctx, name)
}
