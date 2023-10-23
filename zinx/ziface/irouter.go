// @program:     ainx
// @file:        irouter.go
// @author:      ma
// @create:      2023-10-23 16:00
// @description:

// 路由的抽象接口
// 路由里的数据都是IRequest请求

package ziface

type IRouter interface {
	//
	// PreHandle
	// @Description： 在处理Conn业务之前的钩子方法Hook
	// @param request
	//
	PreHandle(request IRequest)
	//
	// Handle
	// @Description: 在处理Conn业务的主方法
	// @param request
	//
	Handle(request IRequest)

	//
	// PostHandle
	// @Description: 在处理Conn业务之后的钩子方法Hook
	// @param request
	//
	PostHandle(request IRequest)
}
