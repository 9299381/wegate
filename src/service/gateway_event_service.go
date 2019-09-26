package service

import "github.com/9299381/wego/contracts"

type GatewayEventService struct {
	next contracts.IService
}

func (it *GatewayEventService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *GatewayEventService) Handle(ctx contracts.Context) error {
	//可记录日志,或者sql
	ctx.Log.Info(ctx.Request())
	return it.next.Handle(ctx)
}
