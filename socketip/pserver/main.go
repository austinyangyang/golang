package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"socketip/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {

		// var buf [128]byte
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read file fialed, err:#v", err)
			return

		}

		// recvStr := string(buf[:n])
		// fmt.Println("revice client send data :", recvStr)

		// conn.Write([]byte(recvStr)) //发送数据
		fmt.Println("收到client发来的数据：", msg)

	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("start listen failed error: #v", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed error: #v", err)
			continue
		}

		go process(conn)

	}

}
