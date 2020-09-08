package main

import "github.com/zstu-club/goso"

func main() {
	server, err := goso.NewTCPServer("", 8999)
	if err != nil {
		panic(err)
	}
	server.RunEcho()
}