// @program:     ainx
// @file:        request.go
// @author:      ma
// @create:      2023-10-23 15:52
// @description:

package znet

import (
	"github.com/peter-matc/Ainx/zinx/ziface"
)

//

type Request struct {
	// 已经和客户端建立好的链接
	conn ziface.IConnection
	// 客户端请求的数据
	data []byte
}

func (r Request) GetConnection() ziface.IConnection {

	return r.conn
}

func (r Request) GetData() []byte {

	return r.data
}
