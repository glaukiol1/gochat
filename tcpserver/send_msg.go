package tcpserver

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func SendMessage(message, host string) {
	CONNECT := host
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Write([]byte(strings.ReplaceAll(message, " ", "-") + "\n"))
	bufio.NewReader(c).ReadString('\n')

}
