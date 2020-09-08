package goso

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	Host string
	Port int
	BufferSize int
	Listener Listener
}

type Listener struct {
	net.Listener
}



func NewTCPServer(host string, port int) (*TCPServer,error){
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return &TCPServer{
		Host:     host,
		Port:     port,
		BufferSize: 4096,
		Listener: Listener{lis},
	}, nil
}

// 获取本地信息
func (s *TCPServer) LocalAddr()string{
	return fmt.Sprintf("%v", s.Listener.Addr())
}

// 设置信息
func (s *TCPServer) SetBufferSize(size int){
	s.BufferSize = size
}

// conn处理类型
type ConnHandler func(conn Conn)

// 运行服务
func (s *TCPServer) Run(handleFun ConnHandler){
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleFun(Conn{conn, s.BufferSize})
	}
}

// echo服务
func (s *TCPServer) EchoConn(conn Conn){
	log.Printf("连接到了 %v", conn.RemoteAddr())
	for {
		buf := make([]byte, s.BufferSize)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("断开连接 %v", conn.RemoteAddr())
			return
		}
		if n == 0 {
			return
		}
		log.Printf("[%v] SAY: (%s) LENGTH:[%d]\n", conn.RemoteAddr(), string(buf[:n]), n)
		conn.Write(buf[:n])
	}
}

// 预定义一些服务
func (s *TCPServer) RunEcho(){
	s.Run(s.EchoConn)
}