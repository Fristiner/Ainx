package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println(err)
			}

			go func(conn net.Conn) {

				dp := NewDataPack()
				for {

					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println(err)
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println(err)
						return
					}
					if msgHead.GetMsgLen() > 0 {

						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())

						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err", err)
							return
						}

						fmt.Println("-----> Recv MsgId = ", msg.Id, " dataLen = ", msg.DataLen, " data = ", string(msg.Data))
					}

				}

			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err: ", err)
		return
	}

	dp := NewDataPack()
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error ", err)
		return
	}

	msg2 := &Message{
		Id:      2,
		DataLen: 6,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error", err)
		return
	}

	sendData := append(sendData1, sendData2...)

	_, err = conn.Write(sendData)
	if err != nil {
		fmt.Println("client Write error ", err)
		return
	}

	select {}
}
