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

	AddRouter(msgID uint32, router IRouter)
	//
	// GetConnMgr
	// @Description: 获取当前server的连接管理器
	// @return IConnManager
	//
	GetConnMgr() IConnManager
	//
	// SetOnConnStart
	// @Description:  设置OnConnStart的钩子函数的方法
	// @param func(connection IConnection)
	//
	SetOnConnStart(func(connection IConnection))
	//
	// SetOnConnStop
	// @Description: 设置OnConnStop的钩子函数的方法
	// @param func(connection IConnection)
	//
	SetOnConnStop(func(connection IConnection))
	//
	// CallOnConnStart
	// @Description: 调用OnConnStart构子函数的方法
	// @param connection
	//
	CallOnConnStart(connection IConnection)
	//
	// CallOnConnStop
	// @Description: 调用OnConnStop构子函数的方法
	// @param connection
	//
	CallOnConnStop(connection IConnection)
}
