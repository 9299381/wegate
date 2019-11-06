package controller

import "github.com/9299381/wego/contracts"

type GatewayEventController struct {
}

func (s *GatewayEventController) Handle(ctx contracts.Context) (interface{}, error) {
	//可记录日志,或者sql
	ctx.Log.Info(ctx.Get("request"))
	return nil, nil

}
