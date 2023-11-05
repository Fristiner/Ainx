// @program:     ainx
// @file:        server.go
// @author:      ma
// @create:      2023-10-23 11:08
// @description:

package znet

import (
	"fmt"
	"net"

	"github.com/peter-matc/Ainx/zinx/utils"
	"github.com/peter-matc/Ainx/zinx/ziface"
)

// Server
// @Description: IServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	// 当前的Server添加一个router，server注册的链接处理对应的业务
	//Router ziface.IRouter
	// 当前server 消息管理模块，用来绑定MsgID
	MsgHandler ziface.IMsgHandle
	//该Server的链接管理器
	ConnMgr ziface.IConnManager
	// 该Server创建链接之后自动调用的Hook函数-onConnStart
	OnConnStart func(conn ziface.IConnection)
	// 该Server销毁链接之前自动调用的Hook函数-onConnStop
	OnConnStop func(connection ziface.IConnection)
}

// NewServer
// @Description: 初始化Server模块
// @param name
// @return ziface.IServer
//
// 初始化Server模块
func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s

}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	// 添加路由方法
	fmt.Println("Add Router success!")
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, Listenner at IP : %s,Port: %d is starting\n", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)

	fmt.Printf("[Zinx] Version : %s, MaxConn : %d, MaxPackageSize : %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)

	fmt.Printf("[Start] Server Listenner at IP: %s,Port %d, is starting\n", s.IP, s.Port)

	go func() {
		// 0 开启消息队列及worker工作池
		s.MsgHandler.StartWorkerPool()
		// 1.获取一个TCP的Addr
		Addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		// 2.监听服务器的地址
		listen, err := net.ListenTCP(s.IPVersion, Addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "  err", err)
			return
		}
		// 连接成功
		fmt.Println("start Zinx server success ", s.Name, " success, Listening")

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

			//设置最大连接个数的判断，如果超过最大连接的数量，那么关闭
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				// TODO：给客户端响应一个超出最大连接的错误包

				_ = conn.Close()
				continue
			}

			// 客户端已经连接做一个最大512 字节的回显业务发什么给他回复什么
			// 将处理新连接的业务方法 和conn进行绑定，得到链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			// 启动当前的业务处理

			//
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
//func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
//
//	// 回显的业务
//	fmt.Println("[Conn Handle] CallBackToClient ")
//	if _, err := conn.Write(data[:cnt]); err != nil {
//		fmt.Println("write back buf err ", err)
//		return errors.New("CallBackToClient error")
//	}
//	return nil
//}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源，状态或者一些已经开辟的链接信息 进行停止或者回收
	fmt.Println("[STOP] Zinx Server name ", s.Name)
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	//TODO： 做一些启动服务器之后的额外业务

	// 阻塞的状态

	select {}
}
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(connection ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("——————> Call OnConStart() ...")
		s.OnConnStart(connection)
	}
}

func (s *Server) CallOnConnStop(connection ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("——————> Call OnConStop() ...")
		s.OnConnStop(connection)
	}
}
