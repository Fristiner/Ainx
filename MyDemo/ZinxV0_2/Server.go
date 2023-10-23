// @program:     ainx
// @file:        Server.go
// @author:      ma
// @create:      2023-10-23 11:10
// @description:

package main

import (
	"github.com/peter-matc/Ainx/zinx/znet"
)

func main() {
	// 1. 创建一个Server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.2]")

	// 2.启动serve

	s.Serve()

}
