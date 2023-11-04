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
	// 其中每一个TaskQueue可以存放这么多个utils.GlobalObject.MaxWorkerTaskLen
	//
	TaskQueue []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		// 从全局配置中获取
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 1. 给当前的worker对应的channel消息队列开辟空间
		// 第i个worker就用第i个空间
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 每个worker可以最多有utils.GlobalObject.MaxWorkerTaskLen个请求

		// 2. 启动当前的worker，阻塞等待消息从channel传递进来
		go m.StartOneWorker(i, m.TaskQueue[i])

	}
}

// StartOneWorker
// @Description: 启动一个worker工作流程
// @receiver m
func (m *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started ...")
	// 不断的阻塞等待消息对应消息队列的消息
	for {
		select {
		//如果有消息过来，出列的就是一个客户端的Request，执行当前Request所绑定的业务方法
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

// SendMsgToTaskQueue
// @Description: 将消息交给TaskQueue，由worker进行处理
// @receiver m
// @param request
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 1.将消息平均分配给不通过的 worker
	// 根据客户端建立的ConnID来进行分配
	// 平均分配的

	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		" request MsgID = ", request.GetMsgID(),
		" to WorkerID = ", workerID)
	// 2.将消息发送给对应的worker的TaskQueue即可
	m.TaskQueue[workerID] <- request
}
