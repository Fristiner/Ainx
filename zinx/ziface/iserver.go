// @program:     ainx
// @file:        iserver.go
// @author:      ma
// @create:      2023-10-23 11:08
// @description:

package ziface

type IServer interface {
	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()

	//添加router功能 路由功能 供客户端链接处理使用
	AddRouter(router IRouter)
}
