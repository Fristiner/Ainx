package main

import (
	"fmt"

	"github.com/peter-matc/Ainx/zinx/ziface"
	"github.com/peter-matc/Ainx/zinx/znet"
)

// ping test 自定义路由

type PingRouter struct {
	znet.BaseRouter
}

// PreHandle
// @Description: Test PreRouter
// @receiver p
// @param request
//func (p *PingRouter) PreHandle(request ziface.IRequest) {
//	fmt.Println("Call Router PreHandle..")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.. \n"))
//	if err != nil {
//		fmt.Println("Call back before ping error ", err)
//	}
//
//}

// Handle
// @Description: Test Handle
// @receiver p
// @param request
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle..")
	// 先读取 客户端数据 ，再回写数据
	fmt.Println("recv  from client: msgID = ", request.GetMsgID(), " , data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("ping ping ping ...."))
	if err != nil {
		fmt.Println(err)
	}

}

// PostHandle
// @Description: Test PostHandle
// @receiver p
// @param request
//func (p *PingRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("Call Router after Handle..")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
//	if err != nil {
//		fmt.Println("Call back  Post ping error ", err)
//	}
//}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (p *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle..")
	// 先读取 客户端数据 ，再回写数据
	fmt.Println("recv  from client: msgID = ", request.GetMsgID(), " , data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("welcome to Zinx Hello router ...."))
	if err != nil {
		fmt.Println(err)
	}

}

// PostHandle
// @Description: Test PostHandle
// @receiver p
// @param request
//func (p *HelloZinxRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("Call Router after Handle..")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("hello router...\n"))
//	if err != nil {
//		fmt.Println("Call back  Post ping error ", err)
//	}
//}

// 创建链接之后的钩子函数
func DoConnectionBegin(connection ziface.IConnection) {
	fmt.Println("=======> DoConnectionBegin is Called ...")
	err := connection.SendMsg(202, []byte("DoConnection Begin"))
	if err != nil {
		fmt.Println("SendMsg id error")
		return
	}

	// 给当前连接创建一些属性

	fmt.Println("Set conn Name ， hoe ")
	connection.SetProperty("Name", "ma")

	connection.SetProperty("Home", "github.com/peter-matc/")

	connection.SetProperty("blog", "gitbook.com")

}

// 链接断开之前需要执行的函数
func DoConnctionLost(connection ziface.IConnection) {
	//
	fmt.Println("=========> DoConnectionLost is called ...")
	fmt.Println("connId = ", connection.GetConnID(), " is Lost ....")
	property, _ := connection.GetProperty("Name")

	fmt.Println("name = ", property)

}

func main() {

	// 1.创建一个server句柄 使用zinx 的 api
	s := znet.NewServer()

	// 2.注册链接Hook钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnctionLost)

	// 3.给当前zinx框架添加一个自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	// 4.启动server
	s.Serve()
}
