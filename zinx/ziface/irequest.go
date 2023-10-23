// @program:     ainx
// @file:        irequest.go
// @author:      ma
// @create:      2023-10-23 15:52
// @description:

package ziface

// IRequest
// @Description:  把客户端请求的链接信息，和请求的数据 包装到了一个Request中
type IRequest interface {
	//
	// GetConnection
	// @Description: 得到当前连接
	// @return IConnection
	//
	GetConnection() IConnection

	//
	// GetData
	// 得到请求的消息数据
	// @Description: 数据
	// @return []byte
	//
	GetData() []byte
}
