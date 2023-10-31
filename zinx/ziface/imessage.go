package ziface

// 将请求的消息封装到一个Message中， 定义抽象的接口

type IMessage interface {
	//
	// GetMsgId
	//  @Description: 获取消息的ID
	//  @return uint32
	//
	GetMsgId() uint32
	//
	// GetMsgLen
	//  @Description: 获取消息的长度
	//  @return uint32
	//
	GetMsgLen() uint32
	//
	// GetData
	//  @Description: 获取消息的内容
	//  @return []byte
	//
	GetData() []byte
	//
	// SetMsgId
	//  @Description: 设置消息的ID
	//  @param uint32
	//
	SetMsgId(uint32)
	//
	// SetData
	//  @Description: 设置消息的内容
	//  @param []byte
	//
	SetData([]byte)
	//
	// SetDataLen
	//  @Description: 设置数据的长度
	//  @param uint322
	//
	SetDataLen(uint32)
}
