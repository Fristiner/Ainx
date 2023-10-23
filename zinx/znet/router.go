// @program:     ainx
// @file:        router.go
// @author:      ma
// @create:      2023-10-23 16:00
// @description:

package znet

import (
	"github.com/peter-matc/Ainx/zinx/ziface"
)

//

// 定一个BaseRouter
// 实现router 先嵌入BaseRouter基类，然后对这个基类方法进行重写
type BaseRouter struct{}

// 只需要写而已 要是用的时候重写即可

func (b *BaseRouter) PreHandle(request ziface.IRequest) {}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}
