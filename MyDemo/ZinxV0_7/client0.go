package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/peter-matc/Ainx/zinx/znet"
)

func main() {
	fmt.Println("client start")

	time.Sleep(1 * time.Second)
	// 1.直接链接远程服务器，得到一个conn链接

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err , exit")
		return
	}

	for {
		// 发送封包的msg
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.6 client test message")))
		if err != nil {
			fmt.Println("pack error ", err)
			return
		}
		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("client write is error ", err)
			return
		}

		// 服务器回复数据 MsgID

		// 先读取head部分
		binaryHead := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, binaryHead)
		if err != nil {
			fmt.Println("read head error ", err)
			break
		}
		iMessage, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error ", err)
			break
		}

		if iMessage.GetMsgLen() > 0 {
			// msg里面有数据
			msg := iMessage.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			_, err = io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("read mag data error ", err)
				return
			}
			fmt.Println("msg id = ", msg.Id, " msgLen = ", msg.DataLen, " msgData = ", string(msg.Data))
		}

		time.Sleep(1 * time.Second)

	}

}
