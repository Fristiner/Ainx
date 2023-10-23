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

// Test PreRouter
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle..")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.. \n"))
	if err != nil {
		fmt.Println("Call back before ping error ", err)
	}

}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle..")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping...\n"))
	if err != nil {
		fmt.Println("Call back  ping error ", err)
	}

}

// Test PostHandle

func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router after Handle..")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("Call back  Post ping error ", err)
	}
}

func main() {

	// 1.创建一个server句柄 使用zinx 的 api
	s := znet.NewServer("[zinx V0.3]")
	// 2.给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})
	// 3.启动server
	s.Serve()

}
