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
	args.Server = "gateway"
	args.Name = "gateway"

	wego.Provider(&providers.EnvProvider{})
	wego.Provider(&provider.LocalProvider{})
	//服务注册到配置中心
	wego.Provider(&providers.ConsulRegistyProvider{})
	wego.Provider(&providers.LogProvider{})
	wego.Provider(&providers.MysqlProvider{})

	wego.Router("gateway", &router.GateWayRouter{})
	wego.Start()

}
