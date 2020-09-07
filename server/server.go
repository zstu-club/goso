package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lin, err := net.Listen("tcp", ":8999")
	log.Printf("Server start %v", lin.Addr())
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lin.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn){
	log.Printf("连接到了 %v", conn.RemoteAddr())
	for {
		buf := make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("断开连接 %v", conn.RemoteAddr())
			return
		}
		if n == 0{
			return
		}
		log.Printf("[%v] SAY: (%s) LENGTH:[%d]\n", conn.RemoteAddr(), string(buf[:n]), n)
		conn.Write(buf[:n])
	}
}
