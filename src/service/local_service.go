package service

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
)

type LocalService struct {
	next contracts.IService
}

func (it *LocalService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *LocalService) Handle(ctx contracts.Context) error {

	params := make(map[string]interface{})
	params["test_rpc_post"] = "test_rpc_post"
	resp := clients.
		Micro("consul_demo").
		Service("demo.post").
		Params(params).
		Run()

	ctx.Response("local", resp.Data)
	return it.next.Handle(ctx)
}
