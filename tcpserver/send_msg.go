package tcpserver

import (
	"bufio"
	"fmt"
	"net"
)

func SendMessage(message, host string) {
	CONNECT := host
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Write([]byte(message + "\n"))
	bufio.NewReader(c).ReadString('\n')

}
