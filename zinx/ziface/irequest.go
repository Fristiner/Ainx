// @program:     ainx
// @file:        irequest.go
// @author:      ma
// @create:      2023-10-23 15:52
// @description:

package ziface

// IRequest 接口
// 把客户端请求的链接信息，和请求的数据 包装到了一个Request中

type IRequest interface {
	// 得到当前链接
	GetConnection() IConnection

	// 得到请求的消息数据
	GetData() []byte
}
