package ziface

/*
消息管理抽象层
*/

type IMsgHandle interface {
	//
	// DoMsgHandler
	//  @Description: 调度｜执行 对应的Router消息处理方法
	//  @param request
	//
	DoMsgHandler(request IRequest)
	//
	// AddRouter
	//  @Description:  为消息添加具体的处理逻辑
	//  @param msgID
	//  @param router
	//
	AddRouter(msgID uint32, router IRouter)

	//
	// StartWorkerPool
	// @Description: 启动worker工作池
	//
	StartWorkerPool()
}
