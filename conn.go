package goso

import "net"

type Conn struct {
	net.Conn
	BufferSize int
}

// 发送数据
func (c *Conn) SendBytes(data []byte) (int, error) {
	return c.Write(data)
}

func (c *Conn) SendString(data string) (int, error) {
	return c.Write([]byte(data))
}

// 接收数据 todo 如何接收全部的数据而不是1024字节
func (c *Conn) RecvBytes() ([]byte, error) {
	buf := make([]byte, c.BufferSize)
	_, err := c.Read(buf)
	return buf, err
}

func (c *Conn) RecvString() (string, error) {
	data, err := c.RecvBytes()
	return string(data), err
}