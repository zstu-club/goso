package goso

import (
	"encoding/json"
	"errors"
	"net"
	"reflect"
	"strconv"
)

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

func (c *Conn) Send(data interface{})(int, error){
	var msg string
	switch value := reflect.ValueOf(data); value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		msg = strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		msg = strconv.FormatUint(value.Uint(), 10)
	case reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
		marshal, err := json.Marshal(value.Interface())
		if err != nil {
			return 0, err
		}
		msg = string(marshal)
	case reflect.Ptr:
		return c.Send(value.Elem().Interface())
	case reflect.String:
		msg = value.String()
	default:
		return 0, errors.New("unsupported type")
	}

	return c.SendString(msg)
}

// 接收数据 todo 如何接收全部的数据而不是1024字节
func (c *Conn) RecvBytes() ([]byte, error) {
	buf := make([]byte, c.BufferSize)
	n, err := c.Read(buf)
	return buf[:n], err
}

func (c *Conn) RecvString() (string, error) {
	data, err := c.RecvBytes()
	return string(data), err
}