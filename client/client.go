package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	// 键盘输入
	input := bufio.NewReader(os.Stdin)

	for {
		buf := make([]byte, 1024)
		n, err := input.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		str := strings.TrimSpace(string(buf[:n]))
		conn.Write([]byte(str))
		n, _ = conn.Read(buf)
		fmt.Println("收到", string(buf[:n]))
	}
}
