// @program:     ainx
// @file:        connection.go
// @author:      ma
// @create:      2023-10-23 11:27
// @description:

package znet

import (
	"fmt"
	"net"

	"github.com/peter-matc/Ainx/zinx/ziface"
)

// 连接模块
type Connection struct {
	// 当前的Socket TCP套接字
	Conn *net.TCPConn
	// 链接的ID
	ConnID uint32
	// 当前的链接状态
	isClosed bool
	// 当前连接绑定的处理业务的API
	HandleApi ziface.HandleFunc

	// 告知当前连接已经退出的channel

	ExitChan chan bool
}

func (c *Connection) StartReader() {
	//
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Printf("connID= ", c.ConnID, " Reader is exit,remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()
	// 当前的处理业务
	for {
		// 读取客户端数据到buf 中 最大512 字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 调用当前连接绑定的HandleAPI
		if err = c.HandleApi(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID ", c.ConnID, " handle is error ", err)
			break
		}

	}

}

func (c *Connection) Start() {
	//TODO implement me
	fmt.Println("Conn Start() ... ConnID = ", c.ConnID)
	// 启动从当前连接的读数据的业务
	go c.StartReader()

	// TODO：启动从当前连接写数据的业务

}

func (c *Connection) Stop() {

	fmt.Println("Conn Stop().. ConnID= ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true
	// 关闭socket 连接

	err := c.Conn.Close()
	if err != nil {
		return
	}
	close(c.ExitChan)

}

func (c *Connection) GetTCPConnection() *net.TCPConn {

	return c.Conn
}

func (c *Connection) GetConnID() uint32 {

	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {

	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {

	return nil
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false, //开启状态
		HandleApi: callback_api,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
