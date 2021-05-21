package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read file fialed, err:#v", err)
			break

		}
		recvStr := string(buf[:n])
		fmt.Println("revice client send data :", recvStr)

		conn.Write([]byte(recvStr)) //发送数据

	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("start listen failed error: #v", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed error: #v", err)
			return
		}

		go process(conn)

	}

}
