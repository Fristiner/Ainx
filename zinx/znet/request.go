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

// Request
// @Description: Request请求体
type Request struct {
	// 已经和客户端建立好的链接
	conn ziface.IConnection
	// 客户端请求的数据
	msg ziface.IMessage
}

func (r Request) GetData() []byte {
	return r.msg.GetData()
}

func (r Request) GetConnection() ziface.IConnection {

	return r.conn
}

func (r Request) GetMsg() ziface.IMessage {

	return r.msg
}

func (r Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
