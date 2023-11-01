package znet

import (
	"fmt"
	"strconv"

	"github.com/peter-matc/Ainx/zinx/ziface"
)

/*
	消息处理模块的实现
*/

type MsgHandle struct {
	//
	//  Apis
	//  @Description: 存放每个MsgId所对应的处理方法
	//
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	// 从Request 中找到msgID

	handle, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID())
	}

	// 根据MsgID 调度对应的业务
	handle.PreHandle(request)
	handle.Handle(request)
	handle.PostHandle(request)
}

func (m *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {

	// 判断当前msg绑定的api处理方法是否已经存在
	if _, ok := m.Apis[msgID]; ok {

		// ID已经注册
		panic("repeat api ,msgID = " + strconv.Itoa(int(msgID)))
	}

	// 2. 添加msg与api的绑定关系
	m.Apis[msgID] = router
	fmt.Println("Add api MsgID = ", msgID, " success!")

}
