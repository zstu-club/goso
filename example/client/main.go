package main

import (
	"fmt"
	"github.com/zstu-club/goso"
)

func main() {
	client, err := goso.NewTCPClient("127.0.0.1", 8999)
	if err != nil {
		panic(err)
	}

	for {
		data := client.InputStrTrimSpace("input: ")
		if len(data) > 0 {
			client.SendString(data)
			recv, err := client.RecvString()
			if err != nil {
				return
			}
			fmt.Println("recv :", recv)
		} else {
			client.Send(&map[string]interface {}{
				"name": "kainhuck",
				"age": 10,
				"hoby": []string{"golang", "python"},
			})
			recv, err := client.RecvString()
			if err != nil {
				return
			}
			fmt.Println("recv :", recv)
		}
	}
}
