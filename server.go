package main

import (
	"fmt"
	"net"

	"github.com/glaukiol1/gochat/tcpserver"
)

func main() {
	PORT := ":" + fmt.Sprint(8777)
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go tcpserver.HandleNewMsg(c)
	}
}
