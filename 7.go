/*
尝试go web
*/
package main

import (
	"fmt"
)

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器启动")
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}
