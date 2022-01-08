package tcpserver

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleNewMsg(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := strings.TrimSpace(string(netData))
		println(c.RemoteAddr().String() + ": " + msg)
		break
	}
}
