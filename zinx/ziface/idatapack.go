package ziface

// 封包 和 拆包 模块
//TODO：拆包过程的实现
// 拆包 先读取头 得到包的内容长度  和 包的类型｜id
// 在从内容读取

//TODO：封包过程实现
// 封包 写message的长度
// 写messageID
// 写message的内容

type IDataPack interface {
	//
	// GetHeadLen
	//  @Description: 获取包的头的长度
	//  @return uint32
	//
	GetHeadLen() uint32
	//
	// Pack
	//  @Description: 封包方法
	//  @param msg
	//  @return []byte
	//  @return error
	//
	Pack(msg IMessage) ([]byte, error)
	//
	// Unpack
	//  @Description: 拆包方法
	//  @param []byte
	//  @return IMessage
	//  @return error
	//
	Unpack([]byte) (IMessage, error)
}
