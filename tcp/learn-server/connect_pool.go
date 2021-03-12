package learn_server

import (
	"fmt"
	"net"
	"time"
)

var Pool = map[string]net.Conn{}

type Server struct {
}

func NewServer() * Server {
	return &Server{}
}

func (s *Server) Start() {
	fmt.Println("Start server...")
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Start error :" , err)
	}

	go Status()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Server accept error :", err)
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	Pool[conn.RemoteAddr().String()] = conn
	for {
		buf := make([]byte, 512)
		l, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			delete(Pool, conn.RemoteAddr().String())
			fmt.Println("移除客户端：", conn.RemoteAddr().String())
			return //终止程序
		}
		fmt.Printf("Received data: %v\n", string(buf[:l]))
	}
}

func Status() {
	tick := time.Tick(10 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("server status")
			for key , v := range Pool {
				fmt.Println("用户：", key, "链接状态：", v.LocalAddr())
			}
		}
	}
}