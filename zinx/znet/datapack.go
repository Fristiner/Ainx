package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/peter-matc/Ainx/zinx/utils"
	"github.com/peter-matc/Ainx/zinx/ziface"
)

// 封包 拆包的模块

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// 四个字节长度
	// DataLen uint32 4 + ID uint32 4
	return 8
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将dataLen写入
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	// 将MsgId 写入
	err1 := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err1 != nil {
		return nil, err1
	}
	// 将data 数据写入
	err2 := binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err2 != nil {
		return nil, err2
	}
	return dataBuff.Bytes(), nil
}

// Unpack
//
//	@Description: 拆包方法  将head信息读出来 之后再根据head信息里的data长度进行读取
//	@receiver dp
//	@param binaryData
//	@return ziface.IMessage
//	@return error
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据的IoReader
	readerBuff := bytes.NewReader(binaryData)
	// 只解压head信息，得到datalen和MsgId
	msg := &Message{}

	err := binary.Read(readerBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}

	// 读取MsgId

	err1 := binary.Read(readerBuff, binary.LittleEndian, &msg.Id)
	if err1 != nil {
		return nil, err1
	}

	// 判断dataLen是否已经超出最大允许长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg data recv! ")
	}

	return msg, nil
}
