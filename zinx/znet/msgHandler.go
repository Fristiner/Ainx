package znet

import (
	"fmt"
	"strconv"

	"github.com/peter-matc/Ainx/zinx/utils"
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
	//
	// WorkerPoolSize
	// @Description: 业务工作worker池的worker数量
	//
	WorkerPoolSize uint32
	//
	// TaskQueue
	// @Description: 负责worker取任务的消息队列
	//
	TaskQueue []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		// 从全局配置中获取
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen),
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

// StartWorkerPool
// @Description:  启动一个worker工作池 开启工作池的任务只能一次
// @receiver m
func (m *MsgHandle) StartWorkerPool() {
	//根据workerPoolSize 分别开始Worker，每个Worker用一个go来承载

}

// StartOneWorker
// @Description: 启动一个worker工作流程
// @receiver m
func (m *MsgHandle) StartOneWorker() {

}
