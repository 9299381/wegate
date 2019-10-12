package controller

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
)

type LocalController struct {
}

func (it *LocalController) Handle(ctx contracts.Context) (interface{}, error) {
	params := make(map[string]interface{})
	params["test_rpc_post"] = "test_rpc_post"
	resp := clients.
		Micro("consul_demo").
		Service("demo.post").
		Params(params).
		Run()

	return resp, nil

}

func (it *LocalController) Valid(ctx contracts.Context) error {
	return nil
}
