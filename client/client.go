package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type TCPClient struct {
	Host string
	Port int
	Conn Conn
	InputReader *bufio.Reader
}

type Conn struct {
	net.Conn
}

func NewTCPClient(host string, port int)(*TCPClient,error){
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return &TCPClient{
		Host: host,
		Port: port,
		Conn: Conn{conn},
		InputReader: bufio.NewReader(os.Stdin),
	}, nil
}

// 从键盘读取数据
func (c *TCPClient) InputBytes(prefix string) []byte {
	buf := make([]byte, 1024)
	fmt.Print(prefix)
	n, _ := c.InputReader.Read(buf)

	return buf[:n]
}

func (c *TCPClient) InputString(prefix string) string{
	return string(c.InputBytes(prefix))
}

func (c *TCPClient) InputStrTrimSpace(prefix string) string{
	return strings.TrimSpace(c.InputString(prefix))
}

func (c *TCPClient) InputBytesTrimSpace(prefix string) []byte{
	return []byte(c.InputStrTrimSpace(prefix))
}

// 发送数据
func (c *TCPClient) SendBytes(data []byte) (int,error){
	return c.Conn.Write(data)
}

func (c *TCPClient)SendString(data string)  (int, error){
	return c.Conn.Write([]byte(data))
}

// 接收数据
func (c *TCPClient) RecvBytes() ([]byte,error){
	buf := make([]byte, 1024)
	_, err := c.Conn.Read(buf)
	return buf, err
}

func (c *TCPClient) RecvString() (string, error){
	data, err := c.RecvBytes()
	return string(data), err
}

func main() {
	client, err := NewTCPClient("127.0.0.1", 8999)
	if err != nil {
		panic(err)
	}

	for {
		data := client.InputBytesTrimSpace("请输入>>> ")
		_, _ = client.SendBytes(data)
		recv, _ := client.RecvString()
		fmt.Println("收到", recv)
	}
}
