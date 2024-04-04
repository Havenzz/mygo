package main

import (
	"fmt"
	"mygo/utils"
	"net"
)

type data struct {
	Msg string `json:"msg"`
}

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		return
	}
	defer conn.Close()

	msg := utils.NewMessage(conn)

	err = msg.WriteData(data{Msg: "哈哈"})
	if err != nil {
		fmt.Println("err = ", err)
	}
	message := make([]byte, 1024)
	for {
		n, err := conn.Read(message)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(message[:n]))
	}
}
