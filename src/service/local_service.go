package service

import "github.com/9299381/wego/contracts"

type LocalService struct {
	next contracts.IService
}

func (it *LocalService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *LocalService) Handle(ctx contracts.Context) error {
	ctx.Response("local", "bbb")
	return it.next.Handle(ctx)
}
