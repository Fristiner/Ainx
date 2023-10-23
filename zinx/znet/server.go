// @program:     ainx
// @file:        server.go
// @author:      ma
// @create:      2023-10-23 11:08
// @description:

package znet

import (
	"errors"
	"fmt"
	"net"

	"github.com/peter-matc/Ainx/zinx/ziface"
)

// IServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP: %s,Port %d, is starting\n", s.IP, s.Port)
	go func() {

		// 1.获取一个TCP的Addr
		Addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error: ", err)
			return
		}

		// 2.监听服务器的地址
		listen, err := net.ListenTCP(s.IPVersion, Addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "  err", err)
			return
		}
		// 连接成功
		fmt.Println("start Zinx server success ", s.Name, " success, Listenning")

		// 分配id

		var cid uint32

		cid = 0

		// 3.阻塞的等待客户端连接，处理客户端链接业务（读写）
		for {
			//如果有客户端连接过来，阻塞会返回
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 客户端已经连接 做一个 最大512字节的回显业务  发什么 给他回复什么

			// 将处理新连接的业务方法 和conn进行绑定，得到链接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			// 启动当前的业务处理

			go dealConn.Start()

			//go func() {
			//	//
			//	for {
			//		buf := make([]byte, 512)
			//
			//		cnt, err2 := conn.Read(buf)
			//		if err2 != nil {
			//			fmt.Println("recv buf err ", err)
			//			continue
			//		}
			//		fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
			//
			//		// 回显功能
			//		if _, err := conn.Write(buf[0:cnt]); err != nil {
			//			fmt.Println("writer back buf err  ", err)
			//			continue
			//		}
			//	}
			//
			//}()
		}
	}()
}

// 定义当前客户端连接所绑定的handle api 后面应该自定义
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {

	// 回显的业务
	fmt.Println("[Conn Handle] CallBackToClient ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源，状态或者一些已经开辟的链接信息 进行停止或者回收
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	//TODO： 做一些启动服务器之后的额外业务

	// 阻塞的状态

	select {}
}
