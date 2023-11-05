// @program:     ainx
// @file:        iconnmanager.go
// @author:      ma
// @create:      2023-11-05 13:45
// @description:

package ziface

//IConnection

type IConnManager interface {
	//
	// Add
	// @Description: 添加链接
	// @param connection
	//
	// 添加链接
	Add(connection IConnection)
	//
	// Remove
	// @Description: 删除链接
	// @param connection
	//
	// 删除链接
	Remove(connection IConnection)
	//
	// Get
	// @Description: 通过connID获取链接
	// @param connID
	// @return IConnection
	// @return error
	//
	Get(connID uint32) (IConnection, error)
	//
	// Len
	// @Description: 得到当前链接总数
	// @return int
	//
	Len() int

	//
	// ClearConn
	// @Description: 清除并终止所有的链接
	//
	ClearConn()
}
