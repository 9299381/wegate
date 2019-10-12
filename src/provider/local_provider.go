package provider

import (
	"github.com/9299381/wegate/src/controller"
	"github.com/9299381/wego"
	"github.com/9299381/wego/filters"
)

type LocalProvider struct {
}

func (it *LocalProvider) Boot() {

}

func (it *LocalProvider) Register() {

	wego.Handler(
		"local",
		filters.New(&controller.LocalController{}))

	wego.Handler(
		"GATEWAY_EVENT_HANDLER",
		filters.New(&controller.GatewayEventController{}))

	//无本地service,只跑本地filter
	endpoint := filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.LimitEndpoint{},
		&filters.JwtEndpoint{},
		&filters.GateWayEndpoint{},
	)
	wego.Handler("limit_remote", endpoint)

}
