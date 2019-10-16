package controller

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
)

type LocalController struct {
	*contracts.Controller
}

func (s *LocalController) Handle(ctx contracts.Context) (interface{}, error) {
	params := make(map[string]interface{})
	params["test_rpc_post"] = "test_rpc_post"
	resp := clients.
		Micro("consul_demo").
		Service("demo.post").
		Params(params).
		Run()

	return resp, nil

}
