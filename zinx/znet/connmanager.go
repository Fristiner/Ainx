// @program:     ainx
// @file:        connmanager.go
// @author:      ma
// @create:      2023-11-05 13:45
// @description:

package znet

import (
	"errors"
	"fmt"
	"sync"

	"github.com/peter-matc/Ainx/zinx/ziface"
)

// 链接管理模块

type ConnManager struct {
	//
	// connections
	// @Description: 管理的链接集合
	//
	connections map[uint32]ziface.IConnection
	//
	// connLock
	// @Description: 保护链接的读写锁
	//
	connLock sync.RWMutex
}

func (c *ConnManager) Add(connection ziface.IConnection) {

	// 保护共享资源map 加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()
	// 将conn加入到ConnManager中

	c.connections[connection.GetConnID()] = connection

	fmt.Println("connection add to ConnManager success: conn num = ", c.Len())

}

func (c *ConnManager) Remove(connection ziface.IConnection) {

	c.connLock.Lock()
	defer c.connLock.Unlock()

	// 删除链接信息
	delete(c.connections, connection.GetConnID())
	fmt.Println("connID = ", connection.GetConnID(), "remove to ConnManager successfully conn num = ", c.Len())

}

func (c *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//TODO implement me
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		// 找到了
		return conn, nil
	} else {
		return nil, errors.New("connection not FOUND!")
	}

}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	// 删除所有的链接
	c.connLock.Lock()
	defer c.connLock.Unlock()

	// 删除conn并停止conn的工作
	for connID, conn := range c.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(c.connections, connID)
	}

	fmt.Println("Clear All connections success! conn num = ", c.Len())
}

// NewConnManager
// @Description: 创建初始化
// @return *ConnManager
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
