package service

import "developer.zopsmart.com/go/gofr/pkg/gofr"

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) Greet(ctx *gofr.Context, name string) (string, error) {
	if name == "" {
		return "hello", nil
	}
	return "hello, " + name, nil
}
