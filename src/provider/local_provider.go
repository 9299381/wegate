package provider

import (
	"github.com/9299381/wegate/src/service"
	"github.com/9299381/wego"
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/services"
)

type LocalProvider struct {
}

func (it *LocalProvider) Boot() {

}

func (it *LocalProvider) Register() {

	ss := services.Chain(&service.LocalService{})
	handler := filters.New(ss)

	wego.Handler("local", handler)

	//无本地service,只跑本地filter
	endpoint := filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.LimitEndpoint{},
		&filters.GateWayEndpoint{},
	)
	wego.Handler("limit_remote", endpoint)

}
