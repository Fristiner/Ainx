package znet

type Message struct {
	Id      uint32 // 消息的ID
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(u uint32) {
	m.Id = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) SetDataLen(u uint32) {
	m.DataLen = u
}

func NewMessage(id uint32, dataLen uint32, data []byte) *Message {
	return &Message{Id: id, DataLen: dataLen, Data: data}
}
