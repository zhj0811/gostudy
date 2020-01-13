package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("hello world")

	lner, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listener creat error", err)
	}
	fmt.Println("waiting for client")
	for {
		conn, err := lner.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connection success")
	fmt.Println("client address: ", conn.RemoteAddr())
	buffer := make([]byte, 1024)
	recvLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error", err)
	}
	strBuffer := string(buffer[:recvLen])
	fmt.Println("Message: ", strBuffer)
	fmt.Println("Message len :", recvLen)
	time.Sleep(time.Second * 1)//等一秒钟，可以看出client里面的read函数有阻塞效果
	sendLen, err := conn.Write([]byte("I am server, you message :" + strBuffer))//将client发过来的消息原样发送回去
	if err != nil {
		fmt.Println("send message error", err)
	}
	fmt.Println("send message success")
	fmt.Println("send message len；", sendLen)
}

