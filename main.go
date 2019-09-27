package main

import (
	"github.com/9299381/wegate/src/provider"
	"github.com/9299381/wegate/src/router"
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/providers"
)

func main() {

	args.Registy = "127.0.0.1:8500"
	args.Server = "gateway,event"
	args.Name = "gateway"

	//环境配置
	//服务注册到配置中心
	wego.Provider(&providers.ConsulRegistyProvider{})

	//本地服务注册
	wego.Provider(&provider.LocalProvider{})

	wego.Router("gateway", &router.GateWayRouter{})
	wego.Start()

}
