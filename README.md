
# 注册本身作为gateway server
~~~~
	wego.Router("gateway", &router.GateWayRouter{})
~~~~

# 注册自身调节handler /gateway/*
~~~~
type GateWayRouter struct {
	*servers.GateWayCommServer
}

func (it *GateWayRouter) Boot() {
	it.GateWayCommServer = servers.NewGateWayCommServer()
}
func (it *GateWayRouter) Register() {
    //本地有服务
	it.Route("GET", "/local", wego.Handler("local"))

	//consul_demo 微服务上提供的 /demo/two 受限控制
	//本地无服务
	it.Route(
		"GET",
		"/consul_demo/demo/two",
		wego.Handler("limit_remote"),
	)

}
~~~~
# 流程 //todo这里以后加入正则匹配
~~~~
 1 解析http请求,形成请求参数
 2 本地handler中是否有注册,如果本地有则跑本地
    本地无服务,也可以进行路由注册,目的是为了进行限速,认证,等等的处理
    本地无服务的handler最后一个为GateWayEndpoint,它返回GATEWAY消息
 3 如若本地没有注册,或者本地 响应 GATEWAY
 4 进行网关代理处理
 5 从consul中查询服务
    如果是http,则进行反向代理
    如果是grpc 则代理请求
 6 事后发送GATEWAY_EVENT_HANDLER 事件
    如果本地有注册GATEWAY_EVENT_HANDLER事件的处理,
    则可进行事后记录日志等的处理
    
 
~~~~

