package goso

import (
	"fmt"
	"testing"
)

func TestNewTCPClient(t *testing.T) {
	client, err := NewTCPClient("127.0.0.1", 8999)
	if err != nil {
		panic(err)
	}
	local := client.LocalAddr()
	fmt.Println("start client in", local)
	for {
		data := client.InputBytesTrimSpace("请输入>>> ")
		_, _ = client.SendBytes(data)
		recv, _ := client.RecvString()
		fmt.Println("接收到>>>", recv)
	}
}