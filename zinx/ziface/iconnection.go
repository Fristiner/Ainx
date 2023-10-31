// @program:     ainx
// @file:        iconnection.go
// @author:      ma
// @create:      2023-10-23 11:27
// @description:

package ziface

import (
	"net"
)

// 定义链接模块的抽象层

type IConnection interface {
	//
	// Start
	// @Description: Start 启动链接 让当前链接开始工作
	//
	Start()

	// Stop 停止链接  结束当前连接的工作
	Stop()

	// GetTCPConnection 获取当前连接绑定的Socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前模块连接的ID
	GetConnID() uint32

	// GetRemoteAddr 获取远程客户端的TCP状态 IP port
	GetRemoteAddr() net.Addr

	// SendMsg
	//  @Description:
	//  @param data
	//  @return error
	// 发送数据，将数据发送给远程客户端
	SendMsg(msgId uint32, data []byte) error
}

// HandleFunc 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
