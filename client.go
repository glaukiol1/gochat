package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/glaukiol1/gochat/tcpserver"
)

func main() {
	go func() { serverStart() }()
	fmt.Print("\nHost to connect to: ")
	var host string
	fmt.Scanln(&host)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		tcpserver.SendMessage(scanner.Text(), host)
	}
}

func serverStart() {
	PORT := ":" + os.Args[1]
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
