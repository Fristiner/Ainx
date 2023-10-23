// @program:     ainx
// @file:        server.go
// @author:      ma
// @create:      2023-10-23 11:08
// @description:

package ziface

type IServer interface {
	//
	// Start
	// @Description: 启动服务器
	//
	// 启动服务器
	Start()
	//
	// Stop
	// @Description: 停止服务器
	//
	// 停止服务器
	Stop()
	//
	// Serve
	// @Description: 运行服务器
	//

	Serve()
	//
	// AddRouter
	// @Description: 添加router功能 路由功能 供客户端链接处理使用
	// @param router
	//

	AddRouter(router IRouter)
}
