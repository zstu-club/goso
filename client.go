package goso

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type TCPClient struct {
	Host        string
	Port        int
	Conn        Conn
	InputReader *bufio.Reader
}

func NewTCPClient(host string, port int) (*TCPClient, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return &TCPClient{
		Host:        host,
		Port:        port,
		Conn:        Conn{conn, 1024},
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

func (c *TCPClient) InputString(prefix string) string {
	return string(c.InputBytes(prefix))
}

func (c *TCPClient) InputStrTrimSpace(prefix string) string {
	return strings.TrimSpace(c.InputString(prefix))
}

func (c *TCPClient) InputBytesTrimSpace(prefix string) []byte {
	return []byte(c.InputStrTrimSpace(prefix))
}

// 发送数据
func (c *TCPClient) SendBytes(data []byte) (int, error) {
	return c.Conn.SendBytes(data)
}

func (c *TCPClient) SendString(data string) (int, error) {
	return c.Conn.SendString(data)
}

// 接收数据
func (c *TCPClient) RecvBytes() ([]byte, error) {
	return c.Conn.RecvBytes()
}

func (c *TCPClient) RecvString() (string, error) {
	return c.Conn.RecvString()
}

// 获取自身信息
func (c *TCPClient) LocalAddr() string {
	return fmt.Sprintf("%v", c.Conn.LocalAddr())
}

// 设置自身信息
func (c *TCPClient) SetBufferSize(size int) {
	c.Conn.BufferSize = size
}
