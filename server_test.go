package goso

import (
	"fmt"
	"testing"
)

func TestNewTCPServer(t *testing.T) {
	server, err := NewTCPServer("127.0.0.1", 8999)
	if err != nil {
		panic(err)
	}
	fmt.Println("server start in", server.LocalAddr())
	server.RunEcho()
}
